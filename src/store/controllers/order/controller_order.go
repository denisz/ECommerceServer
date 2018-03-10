package order

import (
	"store/controllers/common"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	common.Controller
}

// Формирование заказа
func(p *Controller) CheckoutPOST(c *gin.Context) {

}