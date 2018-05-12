package router

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

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


	bytes, err := p.API.Form.FormsOrder(orderID)
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, "application/pdf", bytes)
}
