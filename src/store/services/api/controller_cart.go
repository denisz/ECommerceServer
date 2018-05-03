package api

import (
	. "store/models"
	"github.com/cznic/mathutil"
	"store/delivery/russiaPost"
	"time"
	"store/utils"
	"store/services/emails"
	"github.com/teris-io/shortid"
	"errors"
	"github.com/asdine/storm/q"
)

var (
	ErrNotSupportedMethod  = errors.New("не поддерживается метод доставки")
	ErrEmptyDeliveryMethod = errors.New("пустой метод")
	ErrCheckoutBan         = errors.New("превышен лимит создания заказа")
	ErrEmptyCart           = errors.New("корзина не указана")
	ErrEmptyDelivery       = errors.New("доставка не указана")
	ErrEmptyAddress        = errors.New("адрес не указан")
	ErrNotEnoughQuantity   = errors.New("недостаточно количество продукта")
)

type ControllerCart struct {
	Controller
}

func CreateInvoice() (string, error) {
	return shortid.Generate()
}

func (p *ControllerCart) GetOrCreateCart(cartID int) *Cart {
	// новая сессия
	if cartID == 0 {
		return &Cart{}
	}

	var cart Cart
	err := p.GetStore().From(NodeNamedCarts).One("ID", cartID, &cart)
	if err != nil {
		return &Cart{}
	}

	return &cart
}

//расчет стоимости доставки корзины
func (p *ControllerCart) GetDeliveryPrice(cart *Cart) (Price, error) {
	if cart.Delivery == nil {
		return 0, ErrEmptyDeliveryMethod
	}

	switch cart.Delivery.Provider {
	case DeliveryProviderRussiaPost:
		token := "MmmDeJqGxRlL2MXX4oZiknt25K5mUFEg"
		login := "denisxy12@hotmail.com"
		password := "2Q2sminvc"
		client := russiaPost.NewClient(login, password, token, true)

		mailType := russiaPost.MailTypeONLINE_PARCEL

		switch cart.Delivery.Method {
		case DeliveryMethodEMC:
			mailType = russiaPost.MailTypePARCEL_CLASS_1
		case DeliveryMethodRapid:
			mailType = russiaPost.MailTypeBUSINESS_COURIER
		case DeliveryMethodStandard:
			return 0, nil //бесплатная доставка
		}

		dimension := cart.DimensionCalculate()

		r := &russiaPost.DestinationRequest{
			Mass:          cart.WeightCalculate(),
			IndexFrom:     "430005",
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
			DeclareValue:      100,
			WithSimpleNotice:  false,
			WithOrderOfNotice: false,
		}

		res, err := client.Tariff(r)
		if err != nil {
			return 0, err
		}

		return PriceFloor(Price(res.TotalRate) + Price(res.TotalVat)), nil
	case DeliveryProviderBoxberry:
		return 0, nil
	case DeliveryProviderBaikal:
		return 0, nil
	case DeliveryProviderPEC:
		return 0, nil
	case DeliveryProviderNRG:
		return 0, nil
	default:
		return 0, ErrNotSupportedMethod
	}
}

func (p *ControllerCart) BlockCheckout(clientIP string) error {
	//поиск заказов от этого
	matcher := q.And(q.Eq("ClientIP", clientIP), q.Eq("Status", OrderStatusAwaitingPayment))
	//количество заказов которые ждут оплаты
	totalAwaitPayment, err := p.GetStore().From(NodeNamedOrders).
		Select(matcher).
		Count(new(Order))

	if err != nil {
		return err
	}

	if totalAwaitPayment < 2 {
		return nil
	} else {
		return ErrCheckoutBan
	}
}

func (p *ControllerCart) Update(cart *Cart, update CartUpdateRequest) (*Cart, error) {
	var positions []Position
	//позиции
	cart.Positions = appendIfNeeded(cart.Positions, update.ProductSKU)

	for _, v := range cart.Positions {
		//пустые SKU
		if len(v.ProductSKU) == 0 {
			continue
		}
		if v.ProductSKU == update.ProductSKU {
			switch update.Operation {
			//добавление
			case OperationInsert:
				v.Amount = v.Amount + update.Amount
				//обновление
			case OperationUpdate:
				v.Amount = update.Amount
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
	//проверить вес

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
		return nil, err
	}
	return cart, nil
}

func (p *ControllerCart) SetAddress(cart *Cart, address Address) (*Cart, error)  {
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
		return nil, err
	}

	return cart, nil
}

func (p *ControllerCart) SetDelivery(cart *Cart, delivery Delivery) (*Cart, error) {
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
	//расчет доставки
	deliveryPrice, err := p.GetDeliveryPrice(cart)
	if err != nil {
		return nil, err
	}
	//цена за доставку
	cart.DeliveryPrice = deliveryPrice
	//обновить цену
	cart.PriceCalculate()
	//получаем магазин
	db := p.GetStore().From(NodeNamedCarts)
	//сохранить корзину
	err = db.Save(cart)
	//невозможно сохранить
	if err != nil {
		//ошибка
		return nil, err
	}

	return cart, nil
}

func (p *ControllerCart) Checkout(cart *Cart, session *Session) (*Cart, error) {
	//нету корзины
	if cart.ID == 0 {
		return nil, ErrEmptyCart
	}
	//нету адреса
	if cart.Address == nil {
		return nil, ErrEmptyAddress
	}
	//нету доставки
	if cart.Delivery == nil {
		return nil, ErrEmptyDelivery
	}
	//блокировка
	err := p.BlockCheckout(session.ClientIP)
	if err != nil {
		//отправляем письмо с блокировкой
		go utils.SendEmail(utils.CreateBrand(), emails.Ban{
			EmailRecipient: cart.Address.Email,
			NameRecipient:  cart.Address.Name,
		})
		return nil, err
	}

	//каталог
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	//бакеты
	carts := tx.From(NodeNamedCarts)
	orders := tx.From(NodeNamedOrders)
	catalog := tx.From(NodeNamedCatalog)

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
			return nil, err
		}
		//резервируем количество
		product.Quantity = product.Quantity - v.Amount
		//количество отрицательное будем отклонять заказ
		if product.Quantity < 0 {
			return nil, ErrNotEnoughQuantity
		}
		//сохраняем продукт
		err = catalog.Save(&product)
		if err != nil {
			return nil, err
		}
		// фиксируем поизцию
		positions = append(positions, v)
	}
	//фиксируем позиции
	cart.Positions = positions
	//создание счета
	invoice, err := CreateInvoice()
	if err != nil {
		return nil, err
	}
	//расчет доставки
	deliveryPrice, err := p.GetDeliveryPrice(cart)
	if err != nil {
		return nil, err
	}
	//цена за доставку
	cart.DeliveryPrice = deliveryPrice
	//обновить цену
	cart.PriceCalculate()
	//создаем заказ
	order := Order{
		Status:        OrderStatusAwaitingPayment,
		CreatedAt:     time.Now(),
		Positions:     positions,
		Subtotal:      cart.Subtotal,
		Discount:      cart.Discount,
		ProductPrice:  cart.ProductPrice,
		Total:         cart.Total,
		Delivery:      cart.Delivery,
		DeliveryPrice: cart.DeliveryPrice,
		Address:       cart.Address,
		ClientIP:      session.ClientIP,
		ClientPhone:   cart.Address.Phone,
		Invoice:       invoice,
	}
	//сохраняем заказ
	err = orders.Save(&order)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	//завершаем транзакцию
	tx.Commit()
	//отправить письмо на почту если указана
	//p.Emails <- emails.Receipt{ Order: order }
	//p.Notify <- notify.Notify{ Order: order }
	go utils.SendEmail(utils.CreateBrand(), emails.Receipt{Order: order})

	return cart, nil
}
