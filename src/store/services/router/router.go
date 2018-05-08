package router

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "store/models"
	"store/services/api"
)

type Router struct {
	API *api.API
}

func (p *Router) GetPagination(c *gin.Context) Pagination {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}

	return Pagination{ Limit: limit, Offset: offset}
}

func (p *Router) AbortWithError(c *gin.Context, code int, err error) {
	//c.AbortWithError(code, err)
	//panic
	c.JSON(code, gin.H{
		"message": err.Error(),
		"code": code,
	})
}

func (p *Router) JSON(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
}

func NewRouter(API *api.API) *Router {
	return &Router{
		API: API,
	}
}