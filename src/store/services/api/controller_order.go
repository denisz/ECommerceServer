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
)

type ControllerOrder struct {
	Controller
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

//Поиск заказов
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

		matcher := q.Or(q.Eq("Status", filter.Status), q.Eq("Invoice", filter.Invoice))

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
