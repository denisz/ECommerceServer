package models

import "math"

//Подсчет цены со скидкой
func PriceComputer(price int, discount *Discount, amount int) int {
	if discount != nil {
		switch discount.Type {
		case DiscountTypePercentage:
			sale := float64(price) * float64(discount.Amount / 100)
			return PriceFloor(price -  int(math.Floor(sale)) * amount)
		case DiscountTypeFixedAmount:
			return PriceFloor(price - int(discount.Amount) * amount)
		}
	}

	return PriceFloor(price * amount)
}

//Округлем цену
func PriceFloor(price int) int {
	return price - (price % 100) // отбрасываем копейки
}

//попадание в диапазон
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}