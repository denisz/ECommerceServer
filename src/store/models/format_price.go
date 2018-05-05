package models

import "fmt"

func(p Price) Format() string {
	return fmt.Sprintf("%d руб.", p / 100)
}
