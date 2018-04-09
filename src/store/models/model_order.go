package models

import (
	"time"
)

type Status int32

const (
	// Новый заказ
	OrderStatusDraft Status = 0

	// Формированный заказ
	OrderStatusPending Status = 1

	// В обработке
	OrderStatusProcessing Status = 2

	// Закрыт
	OrderStatusClosed Status = 3

	// Отменен
	OrderStatusCanceled Status = 4
)

type (
	// Квитанция
	Receipt struct {
		// Индентификатор
		ID int `storm:"id,increment"`

		// Номер заказа
		OrderID int

		// Ответ
		Response string

		// Поставщик услуг
		Provider string

		// Пользовательская информация
		Payload string
	}

	// Отправка товара
	Shipment struct {
		// Номер для отслеживания
		Tracking string `json:"tracking"`

		// Поставщик услуг
		Provider string `json:"provider"`

		// Meta данные
		Meta string `json:"meta"`
	}

	// Заказ
	Order struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Позиции в заказе
		Positions []Position `json:"positions"`

		// Адрес доставки
		Address *Address `json:"address"`

		// Доставка
		Delivery *Delivery `json:"delivery"`

		// Квитанция об оплате
		Receipt Receipt `json:"-"`

		// Информация о доставке
		Shipping Shipment `json:"shipping"`

		// Скидка
		Discount *Discount `json:"discount"`

		// Статус заказа
		Status Status `json:"status"`

		// Владелец заказа
		UserID int `json:"userID"`

		// Корзина
		CartID int `json:"cartID"`

		// Счёт на оплату
		Invoice string `storm:"unique" json:"invoice"`

		// Цена на товары
		Subtotal int `json:"subtotal"`

		// Общая цена
		Total int `json:"total"`

		// Цена доставки
		DeliveryPrice int `json:"deliveryPrice"`

		// Комментарий заказа
		Comment string `json:"comment"`

		// Время создания
		CreatedAt time.Time `json:"createdAt"`
	}

	// История измения статуса заказа
	History struct {
		// Инфентификатор
		ID int `storm:"id,increment"`

		// Номер заказа
		OrderID int `json:"orderId"`

		// Индентифкатор оператора
		OperatorID int `json:"operatorId"`

		// Комментарий оператора
		Comment string `json:"comment"`

		// Статус
		Status string `json:"status"`
	}
)
