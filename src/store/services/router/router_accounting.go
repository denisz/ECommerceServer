package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asdine/storm"
	"strconv"
	. "store/models"
)

//получить информацию о отчете
func (p *Router) ReportDetailPOST(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		p.AbortWithError(c, http.StatusBadRequest, nil)
		return
	}

	reportID, err := strconv.Atoi(id)
	if err != nil {
		p.AbortWithError(c, http.StatusBadRequest, err)
		return
	}

	report, err := p.API.Accounting.GetReportByID(reportID)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, report)
}

//список отчетов
func (p *Router) ReportListPOST(c *gin.Context) {
	pagination := p.GetPagination(c)
	reports, err := p.API.Accounting.GetAllReports(pagination)

	if err == storm.ErrNotFound {
		p.AbortWithError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, reports)
}

//поиск отчетов
func (p *Router) SearchReportsPOST(c *gin.Context) {
	var filter FilterAccountingReport

	if err := c.ShouldBindJSON(&filter); err == nil {
		pagination := p.GetPagination(c)

		orders, err := p.API.Accounting.SearchReportsWithFilter(filter, pagination)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}
		p.JSON(c, http.StatusOK, orders)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

//Проверка новые отчетов на Google drive
func (p *Router) CheckReportsInGoogleDrv(c *gin.Context) {
	err := p.API.Accounting.CheckReportsInGoogleDrv()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	err = p.API.Accounting.UpdateQuantity()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

//Поиск расходных накладных
func (p *Router) SearchReceivedReports(c *gin.Context) {
	err := p.API.Accounting.CheckDeliveryReports()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	err = p.API.Accounting.UpdateQuantity()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

func (p *Router) UpdateQuantity(c *gin.Context) {
	err := p.API.Accounting.UpdateQuantity()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

func (p *Router) ClearReports(c *gin.Context) {
	err := p.API.Accounting.ClearReports()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	err = p.API.Accounting.UpdateQuantity()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}