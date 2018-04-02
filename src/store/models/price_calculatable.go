package models

import "math"

// обновление цены у продукта
func(p Product) PriceCalculate() {
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
*/
func(p *Cart) PriceCalculate() {
	//имеется скидка на всю корзину
	if p.Discount != nil {
		//подсчитываем общую сумму корзины
		positions := p.Positions
		//сброс цены
		p.Price = 0
		//обновляем позиции
		p.Positions = []Position{}

		for _, position := range positions {
			//запуск просчета для позиции
			position.PriceCalculate()
			//храним позицию
			p.Positions = append(p.Positions, position)
			//складываем цену без скидки
			p.Price = p.Price + position.Price
		}
		//рассчитываем цены со скидкой
		p.Discount.Price = PriceComputer(p.Price, p.Discount, 1)
		//округляем ценник
		p.Price = PriceFloor(p.Price)
	} else {
		positions := p.Positions
		//подсчитываем общую сумму корзины с учетом скидки каждой позиции
		p.Price = 0
		//обновляем позиции
		p.Positions = []Position{}

		for _, position := range positions {
			//запуск просчета для позиции
			position.PriceCalculate()
			//храним позицию
			p.Positions = append(p.Positions, position)
			//если существует скидка у позиции, складываем цену со скидкой
			if position.Discount != nil {
				p.Price = p.Price + position.Discount.Price
			} else {
				//в противном случаи, складываем цену без скидки
				p.Price = p.Price + position.Price
			}
		}

		if InBetween(p.Price, 6000 * 100, 10000 * 100)  {
			p.Discount = &Discount {
				Type: DiscountTypePercentage,
				Amount: 2.5,
			}
			p.PriceCalculate()
			return
		}

		if InBetween(p.Price, 10000 * 100, 20000 * 100)  {
			p.Discount = &Discount {
				Type: DiscountTypePercentage,
				Amount: 5,
			}
			p.PriceCalculate()
			return
		}

		if InBetween(p.Price, 20000 * 100, math.MaxInt32)  {
			p.Discount = &Discount {
				Type: DiscountTypePercentage,
				Amount: 7.5,
			}
			p.PriceCalculate()
			return
		}
	}
}

// обновление цены у позиции корзины
func(p *Position) PriceCalculate() {
	//продукт
	product := p.Product
	//цена позиции (цена продукта * общее количество)
	p.Price = PriceComputer(product.Price, nil, p.Amount)

	if p.Discount != nil {
		//если существует скидка у позиции, расчитываем цену со скидкой
		p.Discount.Price = PriceComputer(p.Price, p.Discount, 1)
	}
}
