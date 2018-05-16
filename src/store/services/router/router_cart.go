package router

import (
	. "store/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rs/zerolog/log"
)

// корзина
func (p *Router) CartIndexPOST(c *gin.Context) {
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.API.Cart.GetOrCreateCart(session.CartID)
	//отправляем корзину
	p.JSON(c, http.StatusOK, cart)
}

// детальная информация корзины
func (p *Router) CartDetailPOST(c *gin.Context) {
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.API.Cart.GetOrCreateCart(session.CartID)
	//отправляем корзину
	p.JSON(c, http.StatusOK, cart)
}

//  очищаем корзину
func (p *Router) CartClearDELETE(c *gin.Context) {
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.API.Cart.GetOrCreateCart(0)
	//Сохраняем корзину в сессии
	session.CartID = cart.ID
	//отправляем сессию
	WriteSessionToResponse(c, session)
	//отправляем корзину
	p.JSON(c, http.StatusOK, cart)
}

// обновление корзины
func (p *Router) CartUpdatePOST(c *gin.Context) {
	var update CartUpdateRequest

	if err := c.ShouldBindJSON(&update); err == nil {
		//тек. сессия
		session := ReadSessionFromRequest(c)
		//корзина
		cart := p.API.Cart.GetOrCreateCart(session.CartID)

		cart, err := p.API.Cart.Update(cart, update)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}

		//Сохраняем корзину в сессии
		session.CartID = cart.ID
		//отправляем сессию
		WriteSessionToResponse(c, session)
		//отправляем корзину
		p.JSON(c, http.StatusOK, cart)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

// обновление адреса
func (p *Router) CartUpdateAddressPOST(c *gin.Context) {
	var address Address
	//адрес из запроса
	if err := c.ShouldBindJSON(&address); err == nil {
		//тек. сессия
		session := ReadSessionFromRequest(c)
		//корзина
		cart := p.API.Cart.GetOrCreateCart(session.CartID)

		cart, err := p.API.Cart.SetAddress(cart, address)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}
		//сохраняем корзину в сессии
		session.CartID = cart.ID
		//отправляем сессию
		WriteSessionToResponse(c, session)
		//отправляем корзину
		p.JSON(c, http.StatusOK, cart)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

// обновление доставки
func (p *Router) CartUpdateDeliveryPOST(c *gin.Context) {
	var delivery Delivery
	//доставка из запроса
	if err := c.ShouldBindJSON(&delivery); err == nil {
		//тек. сессия
		session := ReadSessionFromRequest(c)
		//корзина
		cart := p.API.Cart.GetOrCreateCart(session.CartID)

		cart, err := p.API.Cart.SetDelivery(cart, delivery)
		if err != nil {
			p.AbortWithError(c, http.StatusInternalServerError, err)
			return
		}
		//сохраняем корзину в сессии
		session.CartID = cart.ID
		//отправляем сессию
		WriteSessionToResponse(c, session)
		//отправляем корзину
		p.JSON(c, http.StatusOK, cart)
	} else {
		p.AbortWithError(c, http.StatusBadRequest, err)
	}
}

// формируем заказ
func (p *Router) CartCheckoutPOST(c *gin.Context) {
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.API.Cart.GetOrCreateCart(session.CartID)
	cart, err := p.API.Cart.Checkout(cart, &Session{ ClientIP:c.ClientIP() })
	if err != nil {
		log.Error().Err(err)
		p.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}
	//сохраняем корзину в сессии
	session.CartID = cart.ID
	//отправляем сессию
	WriteSessionToResponse(c, session)
	//отправляем корзину
	p.JSON(c, http.StatusOK, cart)
}
