package router

import (
	"github.com/gin-gonic/gin"
	. "store/models"
	"store/services/api"
	"net/http"
)

type Router struct {
	API *api.API
}

func (p *Router) GetPagination(c *gin.Context) Pagination {
	pagination := Pagination{ Limit: 10, Offset: 0 }
	c.BindQuery(&pagination)
	return pagination
}

func (p *Router) AbortWithError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"message": err.Error(),
		"code": code,
	})
}

// renderer pdf файл
func (p *Router) PDF(c *gin.Context, data []byte) {
	c.Data(http.StatusOK, "application/pdf", data)
}

// renderer json
func (p *Router) JSON(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
}

func NewRouter(API *api.API) *Router {
	return &Router{
		API: API,
	}
}