package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// загрузка каталога продуктов
func (p *Router) LoaderCatalogFromGoogle(c *gin.Context) {
	err := p.API.Loader.CatalogFromGoogle()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

// загрузка рекламных баннеров
func (p *Router) LoaderAdsFromGoogle(c *gin.Context) {
	err := p.API.Loader.AdsFromGoogle()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

// загрузка цен
func (p *Router) LoaderPricesFromGoogle(c *gin.Context) {
	err := p.API.Loader.PriceFromGoogle()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

// загрузка городов для сдек
func (p *Router) LoaderCDEKCityFromGoogle(c *gin.Context) {
	err := p.API.Loader.CDEKCityFromGoogle()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}

// загрузка городов для поты России
func (p *Router) LoaderRussiaPostTimeFromGoogle(c *gin.Context) {
	err := p.API.Loader.RussiaPostFromGoogle()
	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H{})
}