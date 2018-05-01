package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Загрузка каталога продуктов
func (p *Router) LoaderCatalogFromGoogle(c *gin.Context) {
	err := p.API.Loader.CatalogFromGoogle()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

//Загрузка рекламных баннеров
func (p *Router) LoaderAdsFromGoogle(c *gin.Context) {
	err := p.API.Loader.AdsFromGoogle()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}
