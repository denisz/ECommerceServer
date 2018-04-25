package models

import (
	"fmt"
	"github.com/looplab/fsm"
	"testing"
)

func TestFSMOrder(t *testing.T) {
	f := fsm.NewFSM(
		"yellow",
		fsm.Events{
			{Name: "warn", Src: []string{"green", "red"}, Dst: "yellow"},
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
	)

	fmt.Println(f.Current())
	err := f.Event("warn")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Current())

	fmt.Println(fsm.Visualize(f))
}
