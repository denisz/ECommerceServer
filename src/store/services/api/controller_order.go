package api

import (
	. "store/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
	"strconv"
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

func (p *ControllerOrder) OrderListPOST(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil  {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var orders []Order
	err = p.GetStore().From(NodeNamedOrders).All(&orders, storm.Limit(limit), storm.Skip(offset))
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
			Total: total,
			Limit: limit,
			Offset: offset,
		},
	})
}


func (p *ControllerOrder) UpdatePOST() {

}