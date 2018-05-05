package models

import "fmt"

func (p *Discount) Format() string {
	switch p.Type {
	case DiscountTypeFreeShipping:
		return fmt.Sprintf("Бесплатная доставка")
	case DiscountTypeFixedAmount:
		return fmt.Sprintf("%v руб.", p.Amount)
	case DiscountTypePercentage:
		return fmt.Sprintf("%v%%", p.Amount)
	default:
		return "-"
	}
}