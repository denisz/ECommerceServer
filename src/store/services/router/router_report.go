package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (p *Router) ReportBatchGET(c *gin.Context) {
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

	data, err := p.API.Form.ReportBatch(batchID)

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.PDF(c, data)
}
