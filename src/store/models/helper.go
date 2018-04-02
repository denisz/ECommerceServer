package models

import "math"

func PriceComputer(price int, discount *Discount, amount int) int {
	if discount != nil {
		switch discount.Type {
		case DiscountTypePercentage:
			sale := float64(price) * float64(discount.Amount / 100)
			return (price -  int(math.Floor(sale))) * amount
		case DiscountTypeFixedAmount:
			return (price - int(discount.Amount)) * amount
		}
	}

	return price * amount
}
