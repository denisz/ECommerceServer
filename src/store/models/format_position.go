package models

import "fmt"

func(p *Position) Format() string {
	return fmt.Sprintf("%s x %d", p.Product.Name, p.Amount )
}
