package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
	"strconv"
	. "store/models"
)

//получить информацию о заказе
func (p *Router) OrderDetailPOST(c *gin.Context) {
	invoice := c.Param("invoice")

	if len(invoice) == 0 {
		p.AbortWithError(c, http.StatusBadRequest, nil)
		return
	}

	order, err := p.API.Order.GetOrderByInvoice(invoice)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, order)
}

//список заказов
func (p *Router) OrderListPOST(c *gin.Context) {
	pagination := p.GetPagination(c)
	orders, err := p.API.Order.GetAllOrders(pagination)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, orders)
}

//поиск заказов
func (p *Router) SearchOrdersPOST(c *gin.Context) {
	var filter FilterOrder

	if err := c.ShouldBindJSON(&filter); err == nil {
		pagination := p.GetPagination(c)

		orders, err := p.API.Order.SearchOrdersWithFilter(filter, pagination)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}
		p.JSON(c, http.StatusOK, orders)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

//обновить заказ
func (p *Router) OrderUpdatePOST(c *gin.Context) {
	var update OrderUpdateRequest

	if err := c.ShouldBindJSON(&update); err == nil {
		id := c.Param("id")

		if len(id) == 0 {
			p.AbortWithError(c, http.StatusBadRequest, nil)
			return
		}

		orderID, err := strconv.Atoi(id)
		if err != nil {
			p.AbortWithError(c, http.StatusBadRequest, err)
			return
		}

		//загружаем заказ по id
		order, err := p.API.Order.GetOrderByID(orderID)
		//обновляем заказ
		err = p.API.Order.Update(order, update)

		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}

		p.JSON(c, http.StatusOK, &order)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

//удаляем просроченные заказы
func (p *Router) OrderClearExpiredPOST(c *gin.Context) {
	err := p.API.Order.ClearExpiredOrders()
	if err != nil && err != storm.ErrNotFound {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

//создание партии
func (p *Router) CreateBatchPOST(c *gin.Context) {
	var update BatchRequest

	if err := c.ShouldBindJSON(&update); err == nil {
		err := p.API.Order.CreateBatch(update.OrderIDs)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}
		p.JSON(c, http.StatusOK, gin.H{})
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}