package api

import (
	"net/http"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/asdine/storm/q"
	. "store/models"
)

type ControllerCatalog struct {
	Controller
}

func (p *ControllerCatalog) CollectionPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	var collection Collection
	err := p.GetStore().From(NodeNamedCatalog).One("SKU", sku, &collection)

	if err == storm.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, collection)
}

// Список коллекции
func (p *ControllerCatalog) CollectionsPOST(c *gin.Context) {
	var collections []Collection
	err := p.GetStore().From(NodeNamedCatalog).All(&collections)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, PageCollections{
		Content: collections,
		Cursor: Cursor{
			Total: len(collections),
			Limit: len(collections),
			Offset: 0,
		},
	})
}

func (p *ControllerCatalog) ProductsPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

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
	err = p.GetStore().From(NodeNamedCatalog).Find("CollectionSKU", sku, &products, storm.Limit(limit), storm.Skip(offset))

	if err != nil && err != storm.ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	total, err := p.GetStore().From(NodeNamedCatalog).Select(q.Eq("CollectionSKU", sku)).Count(new(Product))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, product := range products {
		product.PriceCalculate()
	}

	//update prices
	c.JSON(http.StatusOK, PageProducts{
		Content: products,
		Cursor: Cursor{
			Total: total,
			Limit: limit,
			Offset: offset,
		},
	})
}

func (p *ControllerCatalog) ProductPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	var product Product
	err := p.GetStore().From(NodeNamedCatalog).One("SKU", sku, &product)

	if err == storm.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.PriceCalculate()
	c.JSON(http.StatusOK, product)
}


func (p *ControllerCatalog) NotationPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	var notation Notation
	err := p.GetStore().From(NodeNamedCatalog).One("SKU", sku, &notation)

	if err == storm.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, notation)
}

