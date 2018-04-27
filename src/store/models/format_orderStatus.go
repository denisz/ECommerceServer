package models


func (p OrderStatus) Format() string {
	switch p {
	case OrderStatusAwaitingPayment:
		return "Принят. Ожидает оплаты"
	case OrderStatusShipped:
		return "Завершен"
	case OrderStatusAwaitingFulfillment:
		return "В обработке"
	case OrderStatusAwaitingPickup:
		return "Сформирован"
	case OrderStatusAwaitingShipment:
		return "Отправлен"
	case OrderStatusDeclined:
		return "Отклонен"
	case OrderStatusRefunded:
		return "Возрат"
	default:
		return ""
	}
}
