package models

import (
	"time"
)

type OrderStatus string

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
	// Ожидании оплаты (Ждем оплаты)
	OrderStatusAwaitingPayment OrderStatus = "awaitingPayment"

	// Оплаченный заказ, ждем результатов (Ждем отправки)
	OrderStatusAwaitingFulfillment OrderStatus = "awaitingFulfillment"

	//Заказ упакован, ждем отправки (Сформирован)
	OrderStatusAwaitingPickup OrderStatus = "awaitingPickup"

	// Отправлен (Отправлен)
	OrderStatusAwaitingShipment OrderStatus = "awaitingShipment"

	// Доставлен
	OrderStatusShipped OrderStatus = "shipped"

	// Отклонен/Отменен
	OrderStatusDeclined OrderStatus = "declined"

	// Возврат
	OrderStatusRefunded OrderStatus = "refunded"
)

type FilterOrderWhere string

const (
	FilterOrderWhereInvoice   FilterOrderWhere = "invoice"
	FilterOrderWhereDate      FilterOrderWhere = "date"
	FilterOrderWhereRangeDate FilterOrderWhere = "range_date"
	FilterOrderWherePhone     FilterOrderWhere = "phone"
)

func (p OrderStatus) String() string {
	return string(p)
}

type (
	// Квитанция
	Receipt struct {
		// Индентификатор
		ID int `storm:"id,increment"`
		// Номер счета
		Invoice string
		// Ответ
		Response string
		// Поставщик услуг
		Provider string
		// Пользовательская информация
		Payload string
	}

	// Отправка товара
	Shipment struct {
		// Поставщик услуг
		Provider DeliveryProvider `json:"provider"`
		// Способ доставки
		Method DeliveryMethod `json:"method"`
		// Время отправки
		Date time.Time `json:"date"`
		// Номер для отслеживания
		TrackingNumber string `json:"trackingNumber"`
		//Внешний индентификатор
		ExternalNumber string `json:"externalNumber"`
		// Цена за доставку
		Price Price `json:"price"`
		// Meta данные
		Meta string `json:"meta"`
	}

	// Заказ
	Order struct {
		// Индентификатор
		ID int `storm:"id,increment=1000" json:"id"`
		// Позиции в заказе
		Positions []Position `json:"positions"`
		// Адрес доставки
		Address *Address `json:"address"`
		// Доставка
		Delivery *Delivery `json:"delivery"`
		// Квитанция об оплате
		Receipt Receipt `json:"-"`
		// Квитанция номер
		ReceiptNumber string `json:"receiptNumber"`
		// Информация о доставке
		Shipment Shipment `json:"shipment"`
		// Трек доставки
		TrackingNumber string `json:"trackingNumber"`
		// Скидка
		Discount *Discount `json:"discount"`
		// Статус заказа
		Status OrderStatus `json:"status"`
		// Владелец заказа
		OwnerID int `json:"ownerID"`
		// IP автора
		ClientIP string `storm:"index" json:"-"`
		// Номер телефона
		ClientPhone string `storm:"index" json:"-"`
		// Корзина
		CartID int `json:"cartID"`
		// Счёт на оплату
		Invoice string `storm:"unique" json:"invoice"`
		// Цена на товары
		Subtotal Price `json:"subtotal"`
		//Цена товаров
		ProductPrice Price `json:"productPrice"`
		// Цена доставки
		DeliveryPrice Price `json:"deliveryPrice"`
		// Общая цена
		Total Price `json:"total"`
		// Комментарий заказа
		Comment string `json:"comment"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
		// Время обновления
		UpdatedAt time.Time `json:"updatedAt"`
	}

	// История измения статуса заказа
	History struct {
		// Инфентификатор
		ID int `storm:"id,increment"`
		// Номер заказа
		OrderID int `storm:"index" json:"orderID"`
		// Индентифкатор оператора
		OperatorID int `json:"operatorID"`
		// Комментарий оператора
		Comment string `json:"comment"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
		// Статус
		Status OrderStatus `json:"status"`
	}

	// Обновление заказов
	OrderUpdateRequest struct {
		// Изменение статуса
		Status OrderStatus `json:"status"`
		// Трек
		TrackingNumber string `json:"trackingNumber"`
		// Квитанция об оплате
		ReceiptNumber string `json:"receiptNumber"`
		// Ооповещение пользователя
		NoticeRecipient bool `json:"noticeRecipient"`
		// Комеентарий
		Comment string `json:"comment"`
	}

	// Создание партии
	BatchRequest struct {
		// Набор id заказов
		OrderIDs []int `json:"ids"`
	}

	// Фильтр поиска
	FilterOrder struct {
		// Где искать?
		Where FilterOrderWhere `json:"where"`
		// Запрос
		Query string `json:"query"`
		// Начальная дата
		StartDate time.Time `json:"start_date"`
		// Конечная дата
		EndDate time.Time `json:"end_date"`
		// Статус
		Status OrderStatus `json:"status"`
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
