package models

import "math"

//Подсчет цены со скидкой
func PriceComputer(price Price, discount *Discount, amount int) Price {
	if discount != nil {
		switch discount.Type {
		case DiscountTypePercentage:
			sale := math.Floor(float64(price) * float64(discount.Amount / 100))
			return PriceFloor(price -  Price(sale) * Price(amount))
		case DiscountTypeFixedAmount:
			return PriceFloor(price - Price(discount.Amount) * Price(amount))
		}
	}

	return PriceFloor(price * Price(amount))
}

//Округлем цену
func PriceFloor(price Price) Price {
	return price - (price % 100) // отбрасываем копейки
}

//попадание в диапазон
func InBetween(i, min, max Price) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}