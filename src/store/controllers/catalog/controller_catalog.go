package catalog

import (
	"store/controllers/common"
	"net/http"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/asdine/storm/q"
)

type Controller struct {
	common.Controller
}

func (p *Controller) CollectionPOST(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var collection Collection
	err = p.GetStoreNode().One("ID", id, &collection)

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
func (p *Controller) CollectionsPOST(c *gin.Context) {
	var collections []Collection
	err := p.GetStoreNode().All(&collections)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, PageCollections{
		Content: collections,
		Cursor: common.Cursor{
			Total: len(collections),
			Limit: len(collections),
			Offset: 0,
		},
	})
}

func (p *Controller) ProductsPOST(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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
	err = p.GetStoreNode().Find("CollectionID", id, &products, storm.Limit(limit), storm.Skip(offset))

	if err != nil && err != storm.ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	total, err := p.GetStoreNode().Select(q.Eq("CollectionID", id)).Count(new(Product))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, PageProducts{
		Content: products,
		Cursor: common.Cursor{
			Total: total,
			Limit: limit,
			Offset: offset,
		},
	})
}

func (p *Controller) ProductsSalesPOST(c *gin.Context) {
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

	err = p.GetStoreNode().Select(q.Not(q.Eq("Discount", nil))).Limit(limit).Skip(offset).Find(&products)

	if err != nil && err != storm.ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	total, err := p.GetStoreNode().Select(q.Not(q.Eq("Discount", nil))).Count(new(Product))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, PageProducts{
		Content: products,
		Cursor: common.Cursor{
			Total: total,
			Limit: limit,
			Offset: offset,
		},
	})
}

func (p *Controller) ProductPOST(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var product Product
	err = p.GetStoreNode().One("ID", id, &product)

	if err == storm.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *Controller) ProductBySKUPOST(c *gin.Context) {
	SKU, err := strconv.Atoi(c.Param("SKU"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var product Product
	err = p.GetStoreNode().One("SKU", SKU, &product)

	if err == storm.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

