package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func(p *Router) AccountMePOST(c *gin.Context) {
	err := p.API.Account.Me()
	if err != nil {
		p.AbortWithError(c, http.StatusForbidden, err)
		return
	}

	p.JSON(c, http.StatusOK, gin.H {"user": "me"})
}
