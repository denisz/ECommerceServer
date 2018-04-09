package api

import (
	. "store/models"
	"strconv"
	"net/http"
	"github.com/asdine/storm/q"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
)

type ControllerSales struct {
	Controller
}


func(p *ControllerSales) IndexPOST(c *gin.Context) {
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

	var products []Product

	err = p.GetStore().From(NodeNamedCatalog).Select(q.Not(q.Eq("Discount", nil))).Limit(limit).Skip(offset).Find(&products)

	if err != nil && err != storm.ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	total, err := p.GetStore().From(NodeNamedCatalog).Select(q.Not(q.Eq("Discount", nil))).Count(new(Product))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, product := range products {
		product.PriceCalculate()
	}

	c.JSON(http.StatusOK, PageProducts{
		Content: products,
		Cursor: Cursor{
			Total: total,
			Limit: limit,
			Offset: offset,
		},
	})
}

/**
	1. Поиск всех акционных товаров
	2. Удалить акционные товары
	3. Выбрать 3 следующих
 */
func(p *ControllerSales) UpdatePOST(c *gin.Context) {

}