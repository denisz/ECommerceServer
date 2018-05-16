package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "store/models"
	"strconv"
	"github.com/asdine/storm"
)


// Список партий
func (p *Router) SearchBatchesPOST(c *gin.Context) {
	var filter FilterBatch

	if err := c.ShouldBindJSON(&filter); err == nil {
		pagination := p.GetPagination(c)

		orders, err := p.API.Batch.SearchBatchesWithFilter(filter, pagination)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}
		p.JSON(c, http.StatusOK, orders)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

// Детальная информацию о партии
func (p *Router) BatchDetailGET(c *gin.Context) {
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

	//загружаем заказ по id
	batch, err := p.API.Batch.GetBatchByID(batchID)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, batch)
}

// Расформировать партию
func (p *Router) BreakBatchDELETE(c *gin.Context) {
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

	batch, err := p.API.Batch.BreakBatch(batchID)

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, batch)
}

//Отпарвляем данные в ОПС
func (p *Router) CheckInBatchGET(c *gin.Context) {
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

	batch, err := p.API.Batch.CheckInBatch(batchID)

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, batch)
}