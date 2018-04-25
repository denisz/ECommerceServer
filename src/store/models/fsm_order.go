package models

import "github.com/looplab/fsm"


type OrderFsmLabel string

const (
	// Оплатить
	OrderFsmLabelPayment OrderFsmLabel = "payment"

	// Отправить
	OrderFsmLabelShipment OrderFsmLabel = "shipment"

	// Вернуть
	OrderFsmLabelRefuse OrderFsmLabel = "refuse"

	// Отмененить
	OrderFsmLabelDecline OrderFsmLabel = "decline"
)

type OrderFsm struct {
	FSM *fsm.FSM
	Order *Order
}


func CreateOrderFsm(order *Order) *OrderFsm{


	return nil
}
