package order

import (
	"time"
	"store/controllers/account"
	"store/controllers/catalog"
)

type Status int32

const (
	// Новый заказ
	OrderStatusDraft      Status = 0
	// Формированный заказ
	OrderStatusPending    Status = 1
	// В обработке
	OrderStatusProcessing Status = 2
	// Закрыт
	OrderStatusClosed     Status = 3
	// Отменен
	OrderStatusCanceled   Status = 4
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

	// Доставка
	Shipment struct {
		// Номер для отслеживания
		Tracking string `json:"tracking"`
		// Поставщик услуг
		Provider string `json:"provider"`
	}

	// Позиция в заказе
	Item struct {
		// Товар
		Product catalog.Product `json:"productID"`
		// количество
		Amount int `json:"amount"`
	}

	// Заказ
	Order struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`
		// Позиции в заказе
		Items []Item `json:"items"`
		// Адрес доставки
		Address account.Address `json:"address"`
		// Квитанция об оплате
		Receipt Receipt `json:"-"`
		// Информация о доставка
		Shipment Shipment `json:"shipment"`
		// Скидка
		Discount catalog.Discount `json:"discount"`
		// Статус заказа
		Status Status `json:"status"`
		// Владелец заказа
		UserID int `json:"userID"`
		// Счёт на оплату
		Invoice int `json:"invoice"`
		// Налога
		TaxPrice int `json:"taxPrice"`
		// Цена на товары с налога
		TotalPrice int `json:"totalPrice"`
		// Цена доставки
		ShippingPrice int `json:"shippingPrice"`
		// Комментарий клиента заказа
		Comment string `json:"comment"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
	}

	// История измения статуса заказа
	History struct {
		// Инфентификатор
		ID         int `storm:"id,increment"`
		// Номер заказа
		OrderID    int
		// Индентифкатор оператора
		OperatorID int
		// Комментарий оператора
		Comment    string
		// Статус
		Status     string
	}
)
