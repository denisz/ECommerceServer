package api

import (
	. "store/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
)

type ControllerOrder struct {
	Controller
}

//получить информацию о заказе
func (p *ControllerOrder) OrderPOST(c *gin.Context) {
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


func (p *ControllerOrder) UpdatePOST() {

}