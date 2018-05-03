package api

import (
	. "store/models"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"time"
	"errors"
	"store/utils"
	"store/services/emails"
	"fmt"
)


var (
	ErrOrderAlwaysDeclined = errors.New("заказ уже расформирован")
)

type ControllerOrder struct {
	Controller
}

//получить информацию о заказе
func (p *ControllerOrder) GetOrderByInvoice(invoice string) (*Order, error){
	var order Order
	err := p.GetStore().From(NodeNamedOrders).One("Invoice", invoice, &order)

	if err == storm.ErrNotFound {
		return nil, err
	}

	return &order, nil
}

//список заказов
func (p *ControllerOrder) GetAllOrders(pagination Pagination) (*PageOrders, error){
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
	case FilterOrderWhereDate:
		beginningDay := filter.Date.Truncate(24 * time.Hour)
		nextDay := filter.Date.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", beginningDay), q.Lte("CreatedAt", nextDay))
	case FilterOrderWhereInvoice:
		matcher = q.And(matcher, q.Eq("Invoice", filter.Query))
	case FilterOrderWherePhone:
		matcher = q.And(matcher, q.Eq("ClientPhone", filter.Query))
	}

	var orders []Order
	err := p.GetStore().From(NodeNamedOrders).
		Select(matcher).
		Limit(pagination.Limit).
		Skip(pagination.Offset).
		OrderBy("CreatedAt").
		Reverse().
		Find(&orders)

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().From(NodeNamedOrders).
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
		return ErrOrderAlwaysDeclined
	}

	//каталог
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

	//сформирован
	if order.Status == OrderStatusAwaitingPickup {
		//формируем заказ в поставщике доставки
	}

	//сохраняем заказ
	err = orders.Save(&order)

	if err != nil {
		return err
	}

	//история
	history := History {
		OrderID: order.ID,
		Comment: update.Comment,
		Status: update.Status,
		CreatedAt: time.Now(),
	}

	//сохраняем историю изменений
	err = chronology.Save(&history)
	if err != nil {
		return err
	}

	if update.NoticeRecipient {
		p.NoticeRecipient(order)
	}

	//завершаем транзакцию
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
	var ordersBucket = p.GetStore().From(NodeNamedOrders)

	err := ordersBucket.
		Select(matcher).
		Find(&orders)

	if err != nil {
		return err
	}

	for _, order := range orders {
		p.Update(order, OrderUpdateRequest{
			Status: OrderStatusDeclined,
			NoticeRecipient: true,
			Comment: fmt.Sprintf("Изменен атоматом в %v", time.Now().Format("02-01-2006 15:04:05")),
		})
	}

	return nil
}

func (p *ControllerOrder) ResetDeclined(order Order) error {
	return nil
}