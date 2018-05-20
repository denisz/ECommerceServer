package api

import (
	"github.com/asdine/storm/q"
	"github.com/asdine/storm"
	"time"
	"store/delivery/russiaPost"
	. "store/models"
	"fmt"
	"strings"
)

type ControllerBatch struct {
	Controller
}

//поиск партии по фильтру
func (p *ControllerBatch) SearchBatchesWithFilter(filter FilterBatch, pagination Pagination) (*PageBatches, error) {
	matcher := q.True()

	// повставщик услуг
	if len(filter.Provider) != 0 {
		matcher = q.And(matcher, q.Eq("Provider", filter.Provider))
	}

	switch filter.Where {
	case FilterOrderWhereRangeDate:
		startDate := filter.StartDate.Truncate(24 * time.Hour)
		endDate := filter.EndDate.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", startDate), q.Lte("CreatedAt", endDate))
	case FilterOrderWhereDate:
		startDate := filter.StartDate.Truncate(24 * time.Hour)
		endDate := filter.StartDate.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", startDate), q.Lte("CreatedAt", endDate))
	}

	var batches []Batch
	err := p.GetStore().
		From(NodeNamedBatches).
		Select(matcher).
		Limit(pagination.Limit).
		Skip(pagination.Offset).
		OrderBy("CreatedAt").
		Reverse().
		Find(&batches)

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().
		From(NodeNamedBatches).
		Select(matcher).
		Count(new(Batch))

	if err != nil {
		return nil, err
	}

	return &PageBatches{
		Content: batches,
		Cursor: Cursor{
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//получить информацию о заказе
func (p *ControllerBatch) GetBatchByID(id int) (*Batch, error) {
	var batch Batch
	err := p.GetStore().
		From(NodeNamedBatches).
		One("ID", id, &batch)

	if err == storm.ErrNotFound {
		return nil, err
	}

	return &batch, nil
}

/**
Расформировать партию
1. Перевести все заказы в сформированые
2. Удалить партию
3.
 */
func (p *ControllerBatch) BreakBatch(id int) (*Batch, error) {
	//магазин
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	batches := tx.From(NodeNamedBatches)
	orders := tx.From(NodeNamedOrders)

	var batch Batch
	err = batches.One("ID", id, &batch)
	if err != nil {
		return nil, err
	}

	//набор внешних заказов
	var externalIDs []string

	//переводим все заказаы в статус сформированные (OrderStatusAwaitingPickup)
	for _, item := range batch.Orders {
		var order Order
		orders.One("ID", item.ID, &order)
		order.Status = OrderStatusAwaitingPickup
		err := orders.Save(&order)
		if err != nil {
			return nil, err
		}

		externalIDs = append(externalIDs, order.Shipment.ExternalNumber)
	}

	if batch.Provider == DeliveryProviderRussiaPost {
		//Удаляем партию на сайте почта Россия
		_, err := russiaPost.DefaultClient.RestoreBacklog(externalIDs)
		if err != nil {
			return nil, err
		}
	}

	//удаляем партию
	err = batches.DeleteStruct(&batch)
	if err != nil {
		return nil, err
	}

	// Завершаем транзакцию
	tx.Commit()

	return &batch, nil
}

// отправляем данные в ОПС
func (p *ControllerBatch) CheckInBatch(id int) (*Batch, error) {
	var batch Batch
	err := p.GetStore().
		From(NodeNamedBatches).
		One("ID", id, &batch)

	if err == storm.ErrNotFound {
		return nil, err
	}

	if batch.Provider == DeliveryProviderRussiaPost {
		if len(batch.PayloadRussiaPost) > 0 {
			if len(batch.PayloadRussiaPost) > 1 {
				fmt.Printf("russia post has several batch %s", strings.Join(batch.PayloadRussiaPost, ","))
			}

			batchName := batch.PayloadRussiaPost[0]
			//Отправляет по e-mail электронную форму Ф103 в ОПС для регистрации.
			_, err := russiaPost.DefaultClient.CheckIn(batchName)
			if err != nil {
				return nil, err
			}
			return &batch, nil
		}
	}


	return &batch, nil
}