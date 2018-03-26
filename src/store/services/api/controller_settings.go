package api

import (
	"github.com/gin-gonic/gin"
	. "store/models"
	"net/http"
)

type ControllerSettings struct {
	Controller
}

func (p *ControllerSettings) Index(c *gin.Context) {
	db := p.GetSettings()
	var settings Settings

	err := db.Get("settings", "754-3010", &settings)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, &settings)
}

