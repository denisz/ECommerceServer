package models

import "fmt"

func(p *Position) ToFormat() string {
	return fmt.Sprintf("%s x %d", p.Product.Name, p.Amount )
}
