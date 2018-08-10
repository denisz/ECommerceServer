package models

import (
	"math"
	"regexp"
	"strconv"
)

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


func percent(token string) (float64, bool, bool) {
	r, _ := regexp.Compile(`^([0-9.]+)([%]?)$`)
	t := r.FindStringSubmatch(token)

	if len(t) == 3 {
		if t[2] == "%" {
			i, err := strconv.Atoi(t[1])
			if err != nil {
				return 0, false, false
			}

			return float64(i), true, true
		}

		i, err := strconv.Atoi(t[1])
		if err != nil {
			return 0, false, false
		}
		return float64(i), false, true
	}

	return 0, false, false
}

func tokenToDeliveryPeriod(token string) DeliveryPeriod {
	r, _ := regexp.Compile(`^([0-9]+) - ([0-9]+)`)
	t := r.FindStringSubmatch(token)

	if len(t) == 3 {
		min, _ := strconv.Atoi(t[1])
		max, _ := strconv.Atoi(t[2])
		return DeliveryPeriod{
			Min: min,
			Max: max,
		}
	}

	return DeliveryPeriod{}
}