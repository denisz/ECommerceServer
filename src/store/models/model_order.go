package models

import "time"

type OrderStatus int32

/**
	Awaiting Payment — customer has completed the checkout process, but payment has yet to be confirmed. Authorize only transactions that are not yet captured have this status.
	Awaiting Fulfillment — customer has completed the checkout process and payment has been confirmed
	Awaiting Shipment — order has been pulled and packaged and is awaiting collection from a shipping provider
	Awaiting Pickup — order has been packaged and is awaiting customer pickup from a seller-specified location
	Partially Shipped — only some items in the order have been shipped, due to some products being pre-order only or other reasons
	Completed — order has been shipped/picked up, and receipt is confirmed; an Authorize only transaction has been captured; client has paid for their digital product, and their file(s) are available for download
	Shipped — order has been shipped, but receipt has not been confirmed; seller has used the Ship Items action. A listing of all orders with a "Shipped" status can be found under the More tab of the View Orders screen.
	Cancelled — seller has canceled an order, due to a stock inconsistency or other reasons. Stock levels will automatically update depending on your Inventory Settings. Cancelling an order will not refund the order.
	Declined — seller has marked the order as declined for lack of manual payment, or other reasons
	Refunded — seller has used the Refund action. A listing of all orders with a "Refunded" status can be found under the More tab of the View Orders screen.
	Disputed — customer has initiated a dispute resolution process for the PayPal transaction that paid for the order
	Verification Required — order on hold while some aspect (e.g. tax-exempt documentation) needs to be manually confirmed. Orders with this status must be updated manually. Capturing funds or other order actions will not automatically update the status of an order marked Verification Required.
	Partially Refunded — seller has partially refunded the order.
 */
const (
	// Ожидании оплаты
	OrderStatusAwaitingPayment OrderStatus = 0

	// Оплаченный заказ ждем результатов
	OrderStatusAwaitingFulfillment OrderStatus = 1

	// Отклонен
	OrderStatusDeclined OrderStatus = 2

	// Отправлен
	OrderStatusAwaitingShipment OrderStatus = 3

	// Доставлен
	OrderStatusShipped OrderStatus = 4

	// Закрыт
	OrderStatusClosed OrderStatus = 5

	// Возврат
	OrderStatusRefunded OrderStatus = 6

	// Выполнен
	OrderStatusCompleted OrderStatus = 7
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
		Status OrderStatus `json:"status"`

		// Владелец заказа
		UserID int `json:"userID"`

		// Корзина
		CartID int `json:"cartID"`

		// Счёт на оплату
		Invoice string `storm:"unique" json:"invoice"`

		// Цена на товары
		Subtotal int `json:"subtotal"`

		//Цена товаров
		ProductPrice int `json:"productPrice"`

		// Цена доставки
		DeliveryPrice int `json:"deliveryPrice"`

		// Общая цена
		Total int `json:"total"`

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

	// Страницы заказов
	PageOrders struct {
		Content []Order `json:"content"`

		// Курсор
		Cursor
	}

	// Страницы заказов
	PageHistory struct {
		Content []History `json:"content"`

		// Курсор
		Cursor
	}


)
