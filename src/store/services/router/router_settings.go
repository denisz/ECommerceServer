package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *Router) SettingsIndex(c *gin.Context) {
	settings, err := p.API.Settings.GetSettings()

	if err != nil {
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	p.JSON(c, http.StatusOK, settings)
}
