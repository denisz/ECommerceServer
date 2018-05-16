package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// распродажа
func(p *Router) SalesIndexPOST(c *gin.Context) {
	pagination := p.GetPagination(c)
	products, err := p.API.Sales.GetProducts(pagination)

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, products)
}
