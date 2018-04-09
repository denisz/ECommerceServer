package models

import "math"

// обновление цены у продукта
func (p Product) PriceCalculate() {
	if p.Discount != nil {
		p.Discount.Price = PriceComputer(p.Price, p.Discount, 1)
	}
}

// обновляем ценник для корзины
// скидка действует только на цены позиций без скидки
/**
	Скидки на корзину:
	С 6 до 10 тыс 2.5%
	С 10 до 20 тыс 5%
	Свыше 20 тыс. 7.5%
	(продукты со скидкой) + (продукты без скидки) * (динамическую кидку)
*/
func (p *Cart) PriceCalculate() {
	priceSale := 0
	priceWithoutSale := 0
	positions := p.Positions

	//сброс цены
	p.Price = 0
	//сброс позиции
	p.Positions = []Position{}
	//сброс скидки
	p.Discount = nil

	for _, position := range positions {
		//запуск просчета для позиции
		position.PriceCalculate()
		//храним позицию
		p.Positions = append(p.Positions, position)

		if position.Discount != nil {
			priceSale = priceSale + position.Discount.Price
		} else {
			priceWithoutSale = priceWithoutSale + position.Price
		}
	}

	if InBetween(priceWithoutSale, 6000*100, 10000*100) {
		p.Discount = &Discount{
			Type:   DiscountTypePercentage,
			Amount: 2.5,
		}
		p.Discount.Price = priceSale + PriceComputer(priceWithoutSale, p.Discount, 1)
	}

	if InBetween(priceWithoutSale, 10000*100, 20000*100) {
		p.Discount = &Discount{
			Type:   DiscountTypePercentage,
			Amount: 5,
		}
		p.Discount.Price = priceSale + PriceComputer(priceWithoutSale, p.Discount, 1)
	}

	if InBetween(priceWithoutSale, 20000*100, math.MaxInt32) {
		p.Discount = &Discount{
			Type:   DiscountTypePercentage,
			Amount: 7.5,
		}
		p.Discount.Price = priceSale + PriceComputer(priceWithoutSale, p.Discount, 1)
	}

	p.Price = priceWithoutSale + priceSale
	p.Total = p.DeliveryPrice + p.Subtotal()
}

//промежуточный итог
func (p *Cart) Subtotal() int {
	if p.Discount != nil {
		return p.Discount.Price
	}
	return p.Price
}

// обновление цены у позиции корзины
func (p *Position) PriceCalculate() {
	// продукт
	product := p.Product
	// цена позиции (цена продукта * общее количество)
	p.Price = PriceComputer(product.Price, nil, p.Amount)

	if p.Discount != nil {
		// если существует скидка у позиции, расчитываем цену со скидкой
		p.Discount.Price = PriceComputer(p.Price, p.Discount, 1)
	}
}
