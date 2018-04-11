package api

import (
	"github.com/gin-gonic/gin"
	. "store/models"
	"github.com/cznic/mathutil"
	"net/http"
	"github.com/oklog/ulid"
	"store/delivery/russiaPost"
	"crypto/rand"
	"time"
	"store/utils"
	"store/services/emails"
)

type ControllerCart struct {
	Controller
}

func CreateInvoice() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).String()
}

func (p *ControllerCart) CreateCart(session *Session) *Cart {
	cardID := session.CardID
	// новая сессия
	if cardID == 0 {
		return &Cart{}
	}

	var cart Cart
	err := p.GetStore().From(NodeNamedCarts).One("ID", cardID, &cart)
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
			mailType = russiaPost.MailTypeEMS_OPTIMAL
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
		var positions []Position
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
			err := p.GetStore().From(NodeNamedCatalog).One("SKU", v.ProductSKU, &product)
			//продукт недоступен
			if err != nil {
				continue
			}
			//количество не должно превышать допустимое значение
			v.Amount = mathutil.Clamp(v.Amount, 0, product.Quantity)
			//пропускаем позиции с 0 количеством
			if v.Amount <= 0 {
				continue
			}
			//сохраняем продукт
			v.Product = &product
			//скидка на позицию
			v.Discount = product.Discount
			//добавляем позицию
			positions = append(positions, v)
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
		cart.Positions = positions
		//обновить цену
		cart.PriceCalculate()
		//получаем магазин
		db := p.GetStore().From(NodeNamedCarts)
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
		db := p.GetStore().From(NodeNamedCarts)
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
			c.AbortWithError(http.StatusBadRequest, nil)
			return
		}
		//цена за доставку
		cart.DeliveryPrice = p.GetDeliveryPrice(cart)
		//обновить цену
		cart.PriceCalculate()
		//получаем магазин
		db := p.GetStore().From(NodeNamedCarts)
		//сохранить корзину
		err := db.Save(cart)
		//невозможно сохранить
		if err != nil {
			//ошибка
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
	//тек. сессия
	session := ReadSessionFromRequest(c)
	//корзина
	cart := p.CreateCart(session)
	//нету корзины
	if cart.ID == 0 {
		c.AbortWithError(http.StatusNotFound, nil)
		return
	}
	//нету адреса
	if cart.Address == nil {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	//нету доставки
	if cart.Delivery == nil {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	//каталог
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	carts := tx.From(NodeNamedCarts)
	orders := tx.From(NodeNamedOrders)
	catalog := tx.From(NodeNamedCatalog)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer tx.Rollback()

	var positions []Position

	//проверка корзины
	for _, v := range cart.Positions {
		//пустые SKU
		if len(v.ProductSKU) == 0 {
			continue
		}
		//пропускаем позиции с 0 количеством
		if v.Amount <= 0 {
			continue
		}
		//загружаем продукт
		var product Product
		err := catalog.One("SKU", v.ProductSKU, &product)
		//продукт недоступен
		if err != nil {
			//Todo: log
			c.AbortWithError(http.StatusBadRequest, nil)
			return
		}
		//резервируем количество
		product.Quantity = product.Quantity - v.Amount
		//количество отрицательное будем отклонять заказ
		if product.Quantity < 0 {
			//Todo: log
			c.AbortWithError(http.StatusBadRequest, nil)
			return
		}
		//сохраняем продукт
		err = catalog.Save(&product)
		if err != nil {
			//Todo: log
			c.AbortWithError(http.StatusBadRequest, nil)
			return
		}

		positions = append(positions, v)
	}
	//фиксируем позиции
	cart.Positions = positions
	//цена за доставку
	cart.DeliveryPrice = p.GetDeliveryPrice(cart)
	//обновить цену
	cart.PriceCalculate()
	//создаем заказ
	order := Order {
		Status:        OrderStatusAwaitingPayment,
		CreatedAt:     time.Now(),
		Positions:     positions,
		Subtotal:      cart.Subtotal,
		Discount:      cart.Discount,
		Total:         cart.Total,
		Delivery:      cart.Delivery,
		DeliveryPrice: cart.DeliveryPrice,
		Address:       cart.Address,
		Invoice:       CreateInvoice(),
	}
	//сохраняем заказ
	err = orders.Save(&order)
	if err != nil {
		//Todo: log
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//номер заказа
	cart.Invoice = order.Invoice
	//сброс корзины
	cart.Positions = []Position{}
	//сбрасываем доставку
	cart.Delivery = nil
	//сбрасываем цена
	cart.DeliveryPrice = 0
	//сохранить корзину
	err = carts.Save(cart)
	//невозможно сохранить
	if err != nil {
		//Todo: log
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//завершаем транзакцию
	tx.Commit()
	//отправить письмо на почту если указана
	//p.Emails <- emails.Receipt{ Order: order }
	//p.Notify <- notify.Notify{ Order: order }
	go utils.SendEmail(utils.CreateBrand(), emails.Receipt{ Order: order })
	//сохраняем корзину в сессии
	session.CardID = cart.ID
	//отправляем сессию
	WriteSessionToResponse(c, session)
	//отправляем корзину
	c.JSON(http.StatusOK, cart)
}
