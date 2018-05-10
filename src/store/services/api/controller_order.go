package api

import (
	. "store/models"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"time"
	"store/utils"
	"store/services/emails"
	"fmt"
	"store/delivery/russiaPost"
)

type ControllerOrder struct {
	Controller
}

//получить информацию о заказе
func (p *ControllerOrder) GetOrderByInvoice(invoice string) (*Order, error) {
	var order Order
	err := p.GetStore().
		From(NodeNamedOrders).
		One("Invoice", invoice, &order)

	if err == storm.ErrNotFound {
		return nil, err
	}

	return &order, nil
}

//список заказов
func (p *ControllerOrder) GetAllOrders(pagination Pagination) (*PageOrders, error) {
	var orders []Order
	err := p.GetStore().From(NodeNamedOrders).
		AllByIndex("ID", &orders, storm.Limit(pagination.Limit), storm.Skip(pagination.Offset), storm.Reverse())
	if err != nil {
		return nil, err
	}

	total, err := p.GetStore().From(NodeNamedOrders).Count(new(Order))
	if err != nil {
		return nil, err
	}

	return &PageOrders{
		Content: orders,
		Cursor: Cursor{
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//поиск заказов
func (p *ControllerOrder) SearchOrdersWithFilter(filter FilterOrder, pagination Pagination) (*PageOrders, error) {
	matcher := q.True()

	if len(filter.Status) != 0 {
		matcher = q.And(matcher, q.Eq("Status", filter.Status))
	}

	switch filter.Where {
	case FilterOrderWhereRangeDate:
		startDate := filter.StartDate.Truncate(24 * time.Hour)
		endDate := filter.EndDate.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", startDate), q.Lte("CreatedAt", endDate))
	case FilterOrderWhereDate:
		beginningDay := filter.StartDate.Truncate(24 * time.Hour)
		nextDay := filter.StartDate.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", beginningDay), q.Lte("CreatedAt", nextDay))
	case FilterOrderWhereInvoice:
		matcher = q.And(matcher, q.Eq("Invoice", filter.Query))
	case FilterOrderWherePhone:
		matcher = q.And(matcher, q.Eq("ClientPhone", filter.Query))
	}

	var orders []Order
	err := p.GetStore().
		From(NodeNamedOrders).
		Select(matcher).
		Limit(pagination.Limit).
		Skip(pagination.Offset).
		OrderBy("CreatedAt").
		Reverse().
		Find(&orders)

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().
		From(NodeNamedOrders).
		Select(matcher).
		Count(new(Order))

	if err != nil {
		return nil, err
	}

	return &PageOrders{
		Content: orders,
		Cursor: Cursor{
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//Заказ
func (p *ControllerOrder) GetOrderByID(id int) (Order, error) {
	var order Order
	err := p.GetStore().
		From(NodeNamedOrders).
		One("ID", id, &order)

	if err != nil {
		return order, err
	}

	return order, nil
}

//обновить заказ
func (p *ControllerOrder) Update(order Order, update OrderUpdateRequest) error {
	//расформированные заказ
	if order.Status == OrderStatusDeclined {
		return ErrORDER_ALWAYS_DECLINED
	}

	//магазин
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	orders := tx.From(NodeNamedOrders)
	catalog := tx.From(NodeNamedCatalog)
	chronology := tx.From(NodeNamedHistory)

	order.Status = update.Status

	if len(update.ReceiptNumber) > 0 {
		order.ReceiptNumber = update.ReceiptNumber
	}

	if len(update.TrackingNumber) > 0 {
		order.TrackingNumber = update.TrackingNumber
	}

	if len(update.Comment) > 0 {
		order.Comment = update.Comment
	}

	//отменили заказ
	if order.Status == OrderStatusDeclined {
		//производим расформирование заказа
		for _, v := range order.Positions {
			//пропускаем позиции с 0 количеством
			if v.Amount <= 0 {
				continue
			}
			//загружаем продукт
			var product Product
			err := catalog.One("SKU", v.ProductSKU, &product)
			//продукт недоступен
			if err != nil {
				return err
			}
			//возвращаем товар
			product.Quantity = product.Quantity + v.Amount
			//сохраняем продукт
			err = catalog.Save(&product)
			if err != nil {
				return err
			}
		}
	}

	// Сформирован
	if order.Status == OrderStatusAwaitingPickup {
		if order.Delivery.Provider == DeliveryProviderRussiaPost {
			if len(order.Shipment.ExternalNumber) == 0 {
				providerEntity, err := CreateOrderInToRussiaPost(&order)
				if err != nil {
					return err
				}
				order.Shipment.Price = PriceFloor(Price(providerEntity.TotalRate + providerEntity.TotalVat))
				order.Shipment.TrackingNumber = providerEntity.Barcode
				order.Shipment.ExternalNumber = fmt.Sprintf("%d", providerEntity.ID)
			}
		}
	}

	// Сохраняем заказ
	err = orders.Save(&order)
	if err != nil {
		return err
	}

	// История
	history := History{
		OrderID:   order.ID,
		Comment:   update.Comment,
		Status:    update.Status,
		CreatedAt: time.Now(),
	}

	// Сохраняем историю изменений
	err = chronology.Save(&history)
	if err != nil {
		return err
	}

	if update.NoticeRecipient {
		p.NoticeRecipient(order)
	}

	// Завершаем транзакцию
	tx.Commit()

	return nil
}

func (p *ControllerOrder) NoticeRecipient(order Order) error {
	switch order.Status {
	case OrderStatusAwaitingFulfillment:
		//получили оплату
		go utils.SendEmail(utils.CreateBrand(), emails.Processing{Order: order})
	case OrderStatusDeclined:
		//отменили
		//расформировать заказ
		go utils.SendEmail(utils.CreateBrand(), emails.Declined{Order: order})
	case OrderStatusShipped:
		// доставили
		//go utils.SendEmail(utils.CreateBrand(), emails.Receipt{Order: order})
	case OrderStatusAwaitingShipment:
		//отправили
		go utils.SendEmail(utils.CreateBrand(), emails.Shipping{Order: order})
	}

	return nil
}

func (p *ControllerOrder) ClearExpiredOrders() error {
	threshold := time.Now().AddDate(0, 0, -1)
	matcher := q.And(q.Eq("Status", OrderStatusAwaitingPayment), q.Lte("CreatedAt", threshold))
	var orders []Order

	err := p.GetStore().
		From(NodeNamedOrders).
		Select(matcher).
		Find(&orders)

	if err != nil {
		return err
	}

	for _, order := range orders {
		p.Update(order, OrderUpdateRequest{
			Status:          OrderStatusDeclined,
			NoticeRecipient: true,
			Comment:         fmt.Sprintf("Изменен атоматом в %v", time.Now().Format("02-01-2006 15:04:05")),
		})
	}

	return nil
}

// создание партии и перевод заказов в статсу отправлены
func (p *ControllerOrder) CreateBatch(orderIDs []int) error {
	//магазин
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	orders := tx.From(NodeNamedOrders)
	batches := tx.From(NodeNamedBatches)

	var jobs []Order
	var externalNumbersIDs []string

	for _, orderID := range orderIDs {
		var order Order
		err := orders.One("ID", orderID, &order)

		if err != nil {
			return err
		}

		if order.Delivery.Provider != DeliveryProviderRussiaPost {
			continue
		}

		if len(order.Shipment.ExternalNumber) == 0 {
			return fmt.Errorf("not found russiaPostOrder")
		}

		order.Status = OrderStatusAwaitingShipment
		externalNumbersIDs = append(externalNumbersIDs, order.Shipment.ExternalNumber)
		jobs = append(jobs, order)
	}

	shipment, err := russiaPost.DefaultClient.Shipment(externalNumbersIDs, time.Now())
	if err != nil {
		return err
	}

	for _, item := range shipment.Batches {
		batch := Batch{
			Provider:          DeliveryProviderRussiaPost,
			PayloadRussiaPost: item,
			CreatedAt:         time.Now(),
		}

		err := batches.Save(&batch)
		if err != nil {
			return err
		}
	}
	//  собрать документы

	return nil
}
