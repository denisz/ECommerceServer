package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
)


//список партий
func (p *Router) SearchBatchesPOST(c *gin.Context) {
	pagination := p.GetPagination(c)
	orders, err := p.API.Batches.GetAllBatches(pagination)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, orders)
}
