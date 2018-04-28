package api

import (
	. "store/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
	"strconv"
	"github.com/asdine/storm/q"
	"store/utils"
	"store/services/emails"
	"time"
)

type ControllerOrder struct {
	Controller
}

//расформировать заказ
func (p *ControllerOrder) BreakOrder(order *Order) error {
	//каталог
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//бакеты
	catalog := tx.From(NodeNamedCatalog)

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
	//завершаем транзакцию
	tx.Commit()
	return nil
}

//получить информацию о заказе
func (p *ControllerOrder) OrderDetailPOST(c *gin.Context) {
	invoice := c.Param("invoice")

	if len(invoice) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	var order Order
	err := p.GetStore().From(NodeNamedOrders).One("Invoice", invoice, &order)

	if err == storm.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, order)
}

//список заказов
func (p *ControllerOrder) OrderListPOST(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var orders []Order
	err = p.GetStore().From(NodeNamedOrders).AllByIndex("ID", &orders, storm.Limit(limit), storm.Skip(offset), storm.Reverse())
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	total, err := p.GetStore().From(NodeNamedOrders).Count(new(Order))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, PageOrders{
		Content: orders,
		Cursor: Cursor{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	})
}

//поиск заказов
func (p *ControllerOrder) SearchOrderPOST(c *gin.Context) {
	var filter FilterOrder

	if err := c.ShouldBindJSON(&filter); err == nil {
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

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
		err = p.GetStore().From(NodeNamedOrders).
			Select(matcher).
			Limit(limit).
			Skip(offset).
			OrderBy("CreatedAt").
			Reverse().
			Find(&orders)

		if err != nil && err != storm.ErrNotFound {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		total, err := p.GetStore().From(NodeNamedOrders).
			Select(matcher).
			Count(new(Order))

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, PageOrders{
			Content: orders,
			Cursor: Cursor{
				Total:  total,
				Limit:  limit,
				Offset: offset,
			},
		})

	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//обновить заказ
func (p *ControllerOrder) UpdatePOST(c *gin.Context) {
	var json OrderUpdateRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		orderID := c.Param("id")

		if len(orderID) == 0 {
			c.AbortWithError(http.StatusBadRequest, nil)
			return
		}

		orderNumber, err := strconv.Atoi(orderID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var order Order
		err = p.GetStore().From(NodeNamedOrders).One("ID", orderNumber, &order)

		if err == storm.ErrNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		order.Status = json.Status
		order.ReceiptNumber = json.ReceiptNumber
		order.TrackingNumber = json.TrackingNumber

		if json.NoticeRecipient {
			switch order.Status {
			case OrderStatusAwaitingFulfillment:
				//получили оплату
				go utils.SendEmail(utils.CreateBrand(), emails.Processing{Order: order})
			case OrderStatusDeclined:
				//отменили
				go utils.SendEmail(utils.CreateBrand(), emails.Declined{Order: order})
			case OrderStatusShipped:
				// доставили
				//go utils.SendEmail(utils.CreateBrand(), emails.Receipt{Order: order})
			case OrderStatusAwaitingShipment:
				//отправили
				go utils.SendEmail(utils.CreateBrand(), emails.Shipping{Order: order})
			}
		}

		err = p.GetStore().From(NodeNamedOrders).Save(&order)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, &order)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}