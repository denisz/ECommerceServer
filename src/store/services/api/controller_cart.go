package api

import (
	"github.com/gin-gonic/gin"
	. "store/models"
	"github.com/cznic/mathutil"
	"net/http"
	"store/delivery/russiaPost"
)

type ControllerCart struct {
	Controller
}

func (p *ControllerCart) CreateCart(session *Session) *Cart {
	cardID := session.CardID
	// новая сессия
	if cardID == 0 {
		return &Cart{}
	}

	var cart Cart
	err := p.GetStore().One("ID", cardID, &cart)
	if err != nil {
		return &Cart{}
	}

	return &cart
}

//расчет стоимости доставки корзины
func (p *ControllerCart) GetDeliveryPrice(cart *Cart) int {
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
			return 0 //бесплатная доставка
		}

		dimension := cart.DimensionCalculate()

		r := &russiaPost.DestinationRequest{
			Mass:          cart.WeightCalculate(),
			IndexFrom:     "200961",
			IndexTo:       cart.Address.PostalCode,
			MailType:      mailType,
			MailCategory:  russiaPost.MailCategoryORDINARY,
			PaymentMethod: russiaPost.PaymentMethodCASHLESS,
			Dimension: russiaPost.Dimension{
				Width:  dimension.Width,
				Height: dimension.Height,
				Length: dimension.Length,
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
		return 0
	case DeliveryProviderBaikal:
		return 0
	case DeliveryProviderPEC:
		return 0
	case DeliveryProviderNRG:
		return 0
	default:
		return -1
	}
}

func (p *ControllerCart) IndexPOST(c *gin.Context) {
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.CreateCart(session)
	//отправляем корзину
	c.JSON(http.StatusOK, cart)
}

func (p *ControllerCart) DetailPOST(c *gin.Context) {
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.CreateCart(session)
	//отправляем корзину
	c.JSON(http.StatusOK, cart)
}

func (p *ControllerCart) UpdatePOST(c *gin.Context) {
	var json UpdateDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		var readyPositions []Position
		//тек. сессия
		session := ReadSessionFromRequest(c)
		//корзина
		cart := p.CreateCart(session)
		//позиции
		cart.Positions = appendIfNeeded(cart.Positions, json.ProductSKU)

		for _, v := range cart.Positions {
			//пустые SKU
			if len(v.ProductSKU) == 0 {
				continue
			}
			if v.ProductSKU == json.ProductSKU {
				switch json.Operation {
				//добавление
				case OperationInsert:
					v.Amount = v.Amount + json.Amount
					//обновление
				case OperationUpdate:
					v.Amount = json.Amount
					//удаление
				case OperationDelete:
					v.Amount = 0
				}
			}
			//пропускаем позиции с 0 количеством
			if v.Amount <= 0 {
				continue
			}
			//загружаем продукт
			var product Product
			err := p.GetCatalog().One("SKU", v.ProductSKU, &product)
			//продукт недоступен
			if err != nil {
				continue
			}
			//количество не должно превышать допустимое значение
			v.Amount = mathutil.Clamp(v.Amount, 0, product.Quantity)
			//сохраняем продукт
			v.Product = &product
			//скидка на позицию
			v.Discount = product.Discount
			//добавляем позицию
			readyPositions = append(readyPositions, v)
		}
		//указываем возможные способы доставки
		cart.DeliveryProviders = []DeliveryProvider{
			DeliveryProviderRussiaPost,
			DeliveryProviderBoxberry,
		}
		//устанавливаем стандартный способ доставки
		cart.Delivery = &Delivery{
			Provider: DeliveryProviderRussiaPost,
			Method:   DeliveryMethodStandard,
		}
		//цена за стандартную доставку
		cart.DeliveryPrice = 0
		//фиксируем позиции
		cart.Positions = readyPositions
		//обновить цену
		cart.PriceCalculate()
		//получаем магазин
		db := p.GetStore()
		//сохранить корзину
		err := db.Save(cart)
		//невозможно сохранить
		if err != nil {
			//отрпавляем ошибку
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		//Сохраняем корзину в сессии
		session.CardID = cart.ID
		//отправляем сессию
		WriteSessionToResponse(c, session)
		//отправляем корзину
		c.JSON(http.StatusOK, cart)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func (p *ControllerCart) UpdateAddressPOST(c *gin.Context) {
	var address Address
	//адрес из запроса
	if err := c.ShouldBindJSON(&address); err == nil {
		//тек. сессия
		session := ReadSessionFromRequest(c)
		//корзина
		cart := p.CreateCart(session)
		//устанавливаем адрес
		cart.Address = &address
		//указываем возможные способы доставки
		cart.DeliveryProviders = []DeliveryProvider{
			DeliveryProviderRussiaPost,
			DeliveryProviderBoxberry,
		}
		//устанавливаем стандартный способ доставки
		cart.Delivery = &Delivery{
			Provider: DeliveryProviderRussiaPost,
			Method:   DeliveryMethodStandard,
		}
		//цена за стандартную доставку
		cart.DeliveryPrice = 0
		//обновить цену
		cart.PriceCalculate()
		//получаем магазин
		db := p.GetStore()
		//сохранить корзину
		err := db.Save(cart)
		//невозможно сохранить
		if err != nil {
			//отрпавляем ошибку
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		//сохраняем корзину в сессии
		session.CardID = cart.ID
		//отправляем сессию
		WriteSessionToResponse(c, session)
		//отправляем корзину
		c.JSON(http.StatusOK, cart)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func (p *ControllerCart) UpdateDeliveryPOST(c *gin.Context) {
	var delivery Delivery
	//доставка из запроса
	if err := c.ShouldBindJSON(&delivery); err == nil {
		//тек. сессия
		session := ReadSessionFromRequest(c)
		//корзина
		cart := p.CreateCart(session)
		//сброс доставки
		cart.Delivery = nil
		//поиск в доступных провайдерах
		for _, provider := range cart.DeliveryProviders {
			if provider == delivery.Provider {
				//установить доставку
				cart.Delivery = &delivery
				if cart.Delivery.Provider == DeliveryProviderRussiaPost && cart.Delivery.Method == "" {
					cart.Delivery.Method = DeliveryMethodStandard
				}
			}
		}
		//если нету доставки
		if cart.Delivery == nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		//цена за доставку
		cart.DeliveryPrice = p.GetDeliveryPrice(cart)
		//обновить цену
		cart.PriceCalculate()
		//получаем магазин
		db := p.GetStore()
		//сохранить корзину
		err := db.Save(cart)
		//невозможно сохранить
		if err != nil {
			//отрпавляем ошибку
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		//сохраняем корзину в сессии
		session.CardID = cart.ID
		//отправляем сессию
		WriteSessionToResponse(c, session)
		//отправляем корзину
		c.JSON(http.StatusOK, cart)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func (p *ControllerCart) CheckoutPOST(c *gin.Context) {

}
