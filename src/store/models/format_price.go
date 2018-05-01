package models

import "fmt"

func(p Price) ToFormat() string {
	return fmt.Sprintf("%d руб.", p / 100)
}
