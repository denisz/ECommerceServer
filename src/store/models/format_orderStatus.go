package models


func (p OrderStatus) Format() string {
	switch p {
	case OrderStatusAwaitingPayment:
		return "Создан. Ожидает оплаты"
	case OrderStatusClosed:
		return "Закрыт"
	case OrderStatusCompleted:
		return "Выполнен"
	case OrderStatusShipped:
		return "Готов к выдаче"
	case OrderStatusAwaitingFulfillment:
		return "В обработке"
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
