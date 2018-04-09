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
	db := p.GetStore().From(NodeNamedSettings)
	var settings Settings

	db.Get("settings", "754-3010", &settings)
	c.JSON(http.StatusOK, &settings)
}

