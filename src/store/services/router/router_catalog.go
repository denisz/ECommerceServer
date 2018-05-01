package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
	. "store/models"
)

// Коллекция
func (p *Router) CatalogCollectionDetailPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	collection, err := p.API.Catalog.GetCollectionBySKU(sku)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, collection)
}

// Список коллекции
func (p *Router) CatalogCollectionsPOST(c *gin.Context) {
	collections, err := p.API.Catalog.GetAllCollections()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, collections)
}

// Список товаров
func (p *Router) CatalogProductsPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		p.AbortWithError(c, http.StatusBadRequest, nil)
		return
	}

	pagination := p.GetPagination(c)

	products, err := p.API.Catalog.GetProductsByCollectionSKU(sku, pagination)
	if err != nil {
		p.AbortWithError(c,http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, products)
}

//Поиск по наименованию товара
func (p *Router) CatalogSearchProductsPOST(c *gin.Context) {
	var filter FilterCatalog

	if err := c.ShouldBindJSON(&filter); err == nil {
		pagination := p.GetPagination(c)

		products, err := p.API.Catalog.SearchProductsWithFilter(filter, pagination)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}

		p.JSON(c, http.StatusOK, products)
	} else {
		p.AbortWithError(c,http.StatusBadRequest, err)
	}
}

//Продукт
func (p *Router) CatalogProductDetailPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	product, err := p.API.Catalog.GetProductBySKU(sku)
	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, product)
}

//Описания продукта
func (p *Router) CatalogNotationPOST(c *gin.Context) {
	sku := c.Param("sku")

	if len(sku) == 0 {
		p.AbortWithError(c, http.StatusBadRequest, nil)
		return
	}

	notation, err := p.API.Catalog.GetNotationBySKU(sku)
	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, notation)
}