package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cznic/mathutil"
	"store/delivery/russiaPost"
	. "store/models"
)

type ControllerCart struct {
	Controller
}

func (p *ControllerCart) IndexPOST(c *gin.Context) {
	session := ReadCartFromRequest(c)
	c.JSON(http.StatusOK, gin.H{
		"Positions": session.Positions,
	})
}

func (p *ControllerCart) GetDetailCart(session *Session) *Cart {
	cart := Cart{
		Address: session.Address,
		Delivery: session.Delivery,
		DeliveryPrice: session.DeliveryPrice,
	}
	for _, v := range session.Positions {
		if len(v.ProductSKU) == 0 || v.Amount <= 0 {
			continue
		}

		var product Product
		err := p.GetCatalog().One("SKU", v.ProductSKU, &product)
		if err != nil {
			continue
		}

		position := Position {
			Product:    product,
			Amount:     v.Amount,
			ProductSKU: v.ProductSKU,
			Discount:   product.Discount,
		}

		cart.Positions = append(cart.Positions, position)
	}

	cart.PriceCalculate()

	return &cart
}

//расчет стоимости доставки корзины
func (p *ControllerCart) GetDeliveryPrice(session *Session) int {
	cart := p.GetDetailCart(session)
	if cart.Delivery == nil {
		return -1
	}

	switch cart.Delivery.Provider {
	case DeliveryProviderRussiaPost:
		token := "9a9mk3FmmR1E84cn7FHMlz9Kjm5NHAC6"
		login := "viktor@otdeldostavok.ru"
		password := "123456qQ"
		client := russiaPost.NewClient(login, password, token, true)

		mailType := russiaPost.MailTypeONLINE_PARCEL

		switch cart.Delivery.Method {
		case DeliveryMethodEMC:
			mailType = russiaPost.MailTypeEMS
		case DeliveryMethodRapid:
			mailType = russiaPost.MailTypeBUSINESS_COURIER
		case DeliveryMethodStandard:
			mailType = russiaPost.MailTypeONLINE_PARCEL
		}

		r := &russiaPost.DestinationRequest{
			Mass:          2000,
			IndexFrom:     "200961",
			IndexTo:       cart.Address.PostalCode,
			MailType:      mailType,
			MailCategory:  russiaPost.MailCategoryORDINARY,
			PaymentMethod: russiaPost.PaymentMethodCASHLESS,
			Dimension: russiaPost.Dimension{
				Width:  10,
				Height: 10,
				Length: 10,
			},
			Fragile:           false,
			DeclareValue:      3000,
			WithSimpleNotice:  false,
			WithOrderOfNotice: false,
		}

		res, err := client.Tariff(r)
		if err != nil {
			return -1
		}

		return PriceFloor(res.TotalRate + res.TotalVat)
	case DeliveryProviderBoxberry:
		return -1
	default:
		return -1
	}
}


// Детальная информация корзины
func (p *ControllerCart) DetailPOST(c *gin.Context) {
	session := ReadCartFromRequest(c)
	c.JSON(http.StatusOK, p.GetDetailCart(session))
}

/// Обновить позицию
func (p *ControllerCart) UpdatePOST(c *gin.Context) {
	var json UpdateDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		var positions []SessionPosition
		session := ReadCartFromRequest(c)
		origPositions := AppendIfNeeded(session.Positions, json.ProductSKU)

		for _, v := range origPositions {
			if v.ProductSKU == json.ProductSKU {
				switch json.Operation {
				case OperationInsert:
					v.Amount = v.Amount + json.Amount
				case OperationUpdate:
					v.Amount = json.Amount
				case OperationDelete:
					v.Amount = 0
				}
			}

			if v.Amount > 0 {
				if len(v.ProductSKU) == 0 {
					continue
				}

				var product Product
				err := p.GetCatalog().One("SKU", v.ProductSKU, &product)
				if err != nil {
					continue
				}
				v.Amount = mathutil.Clamp(v.Amount, 0, product.Quantity)
				positions = append(positions, v)
			}
		}

		session.Positions = positions
		WriteCartToResponse(c, session)
		c.JSON(http.StatusOK, p.GetDetailCart(session))
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//Сохраняем адрес в сессии для будующих покупок
func (p *ControllerCart) UpdateAddressPOST(c *gin.Context) {
	var json Address

	if err := c.ShouldBindJSON(&json); err == nil {
		session := ReadCartFromRequest(c)
		//Добавить валидацию
		session.Address = &json
		WriteCartToResponse(c, session)
		c.JSON(http.StatusOK, p.GetDetailCart(session))
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//Сохраняем выбор доставки
func (p *ControllerCart) UpdateDeliveryPOST(c *gin.Context) {
	var json Delivery

	if err := c.ShouldBindJSON(&json); err == nil {
		session := ReadCartFromRequest(c)
		//Добавить валидацию
		session.Delivery = &json
		session.DeliveryPrice = p.GetDeliveryPrice(session)
		WriteCartToResponse(c, session)
		c.JSON(http.StatusOK, p.GetDetailCart(session))
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func (p *ControllerCart) CheckoutPOST(c *gin.Context) {
	//по session создаем заказ

	//очищаем корзину

	//бронируем количество товара

	//возвращаем информацию о заказе
}