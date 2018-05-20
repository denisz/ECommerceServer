package router

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// бланк заказа
func(p *Router) FormsOrderPOST(c *gin.Context) {
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


	data, err := p.API.Form.FormsOrder(orderID)
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.PDF(c, data)
}


// бланк партии
func (p *Router) FormsBatchPOST(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		p.AbortWithError(c, http.StatusBadRequest, nil)
		return
	}

	batchID, err := strconv.Atoi(id)
	if err != nil {
		p.AbortWithError(c, http.StatusBadRequest, err)
		return
	}

	data, err := p.API.Form.FormsBatch(batchID)
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.PDF(c, data)
}