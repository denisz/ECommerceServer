package catalog

import (
	"store/controllers/common"
	"net/http"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/cznic/mathutil"
)

type Controller struct {
	common.Controller
}


func (p *Controller) CollectionGET(c *gin.Context) {
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
func (p *Controller) CollectionsGET(c *gin.Context) {
	var collections []Collection
	err := p.GetStoreNode().All(&collections)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, PageCollections{
		Content: collections,
		Cursor: common.Cursor{
			Limit: len(collections),
			Offset: 0,
			Last: true,
		},
	})
}

func (p *Controller) ProductsGET(c *gin.Context) {
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
	err = p.GetStoreNode().Find("CollectionID", id, &products, storm.Limit(limit + 1), storm.Skip(offset))

	if err != nil && err != storm.ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, PageProducts{
		Content: products[:mathutil.Max(len(products) - 1, 0)],
		Cursor: common.Cursor{
			Limit: limit,
			Offset: offset,
			Last: len(products) < limit,
		},
	})
}

func (p *Controller) ProductGET(c *gin.Context) {
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