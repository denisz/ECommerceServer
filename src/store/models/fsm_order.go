package models

import (
	"github.com/looplab/fsm"
	"strconv"
	"fmt"
)

type OrderFsm struct {
	FSM *fsm.FSM
	Order *Order
}

func(p *OrderFsm) Event(status OrderStatus, args ...interface{}) error {
	return p.FSM.Event(status.String(), args)
}

func (p *OrderFsm) Current() OrderStatus {
	i, _ := strconv.Atoi(p.FSM.Current())
	return OrderStatus(i)
}

func CreateOrderFsm(order *Order) *OrderFsm {
	f := OrderFsm {
		FSM : fsm.NewFSM(
			order.Status.String(),
			fsm.Events{
				{Name: OrderStatusAwaitingFulfillment.String(), Src: []string{"green", "red"}, Dst: "yellow"},
				{Name: "panic", Src: []string{"yellow"}, Dst: "red"},
				{Name: "panic", Src: []string{"green"}, Dst: "red"},
				{Name: "calm", Src: []string{"red"}, Dst: "yellow"},
				{Name: "clear", Src: []string{"yellow"}, Dst: "green"},
			},
			fsm.Callbacks{
				"before_warn": func(e *fsm.Event) {
					fmt.Println("before_warn")
				},
				"before_event": func(e *fsm.Event) {
					fmt.Println("before_event")
				},
				"leave_green": func(e *fsm.Event) {
					fmt.Println("leave_green")
				},
				"leave_state": func(e *fsm.Event) {
					fmt.Println("leave_state")
				},
				"enter_yellow": func(e *fsm.Event) {
					fmt.Println("enter_yellow")
				},
				"enter_state": func(e *fsm.Event) {
					fmt.Println("enter_state")
				},
				"after_warn": func(e *fsm.Event) {
					fmt.Println("after_warn")
				},
				"after_event": func(e *fsm.Event) {
					fmt.Println("after_event")
				},
			},
		),
		Order: order,
	}

	return &f
}
