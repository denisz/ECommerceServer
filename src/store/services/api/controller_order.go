package api

import (
	"github.com/gin-gonic/gin"
	. "store/models"
)

type ControllerOrder struct {
	Controller
}

// Формирование заказа
func(p *ControllerOrder) CheckoutPOST(c *gin.Context) {

}