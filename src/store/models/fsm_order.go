package models

import (
	"github.com/looplab/fsm"
	"fmt"
)

type OrderFsm struct {
	FSM   *fsm.FSM
	Order *Order
}

var (
	AllOrderStatuses = []string {
		OrderStatusAwaitingPayment.String(),
		OrderStatusAwaitingFulfillment.String(),
		OrderStatusAwaitingPickup.String(),
		OrderStatusAwaitingShipment.String(),
		OrderStatusShipped.String(),
		OrderStatusDeclined.String(),
		OrderStatusRefunded.String(),
	}
)

func (p *OrderFsm) Event(status OrderStatus, args ...interface{}) error {
	return p.FSM.Event(status.String(), args)
}

func (p *OrderFsm) Current() OrderStatus {
	return OrderStatus(p.FSM.Current())
}

func CreateOrderFsm(order *Order) *OrderFsm {
	f := OrderFsm{
		FSM: fsm.NewFSM(
			order.Status.String(),
			fsm.Events{
				{
					Name: OrderStatusAwaitingFulfillment.String(),
					Src:  AllOrderStatuses,
					Dst:  OrderStatusAwaitingFulfillment.String(),
				},
				{
					Name: OrderStatusAwaitingPickup.String(),
					Src:  AllOrderStatuses,
					Dst:  OrderStatusAwaitingPickup.String(),
				},
				{
					Name: OrderStatusAwaitingShipment.String(),
					Src:  AllOrderStatuses,
					Dst:  OrderStatusAwaitingShipment.String(),
				},
				{
					Name: OrderStatusShipped.String(),
					Src:  AllOrderStatuses,
					Dst:  OrderStatusShipped.String(),
				},
				{
					Name: OrderStatusDeclined.String(),
					Src:  AllOrderStatuses,
					Dst:  OrderStatusDeclined.String(),
				},
				{
					Name: OrderStatusRefunded.String(),
					Src:  AllOrderStatuses,
					Dst:  OrderStatusRefunded.String(),
				},
			},
			fsm.Callbacks{
				fmt.Sprintf("after_%s", OrderStatusAwaitingFulfillment.String()): func(e *fsm.Event) {
					//оплачен
					fmt.Printf("after %s \n", OrderStatusAwaitingFulfillment.String())
				},
				fmt.Sprintf("after_%s", OrderStatusDeclined.String()): func(e *fsm.Event) {
					//отменен
					fmt.Printf("after %s \n", OrderStatusDeclined.String())
				},
				fmt.Sprintf("after_%s", OrderStatusShipped.String()): func(e *fsm.Event) {
					//доставлен
					fmt.Printf("after %s \n", OrderStatusShipped.String())
				},
				fmt.Sprintf("after_%s", OrderStatusAwaitingShipment.String()): func(e *fsm.Event) {
					//отправлен
					fmt.Printf("after %s \n", OrderStatusAwaitingShipment.String())
				},
			},
		),
		Order: order,
	}

	fsm.Visualize(f.FSM)

	return &f
}
