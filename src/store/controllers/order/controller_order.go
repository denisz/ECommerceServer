package order

import (
	"store/controllers/common"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	common.Controller
}

// Формирование и оплата заказа
func(p *Controller) CheckoutPOST(c *gin.Context) {

}

// Отмена заказа пользователем
func(p *Controller) UserCanceledPOST(c *gin.Context) {

}

// Отмена заказа админом
func(p *Controller) AdminCanceledPOST(c *gin.Context) {

}

// Обновление заказа
func(p *Controller) UpdateAddressPOST(c *gin.Context) {

}

// Обновление доставки
func(p *Controller) UpdateShippingPOST(c *gin.Context) {

}

// Обновление статуса заказа
func(p *Controller) AdminUpdateStatePOST(c *gin.Context) {

}