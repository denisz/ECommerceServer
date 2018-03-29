package models


// обновление цены у продукта
func(p Product) PriceCalculate() {
	if p.Discount != nil {
		p.Discount.Price = PriceComputer(p.Price, p.Discount, 1)
	}
}


// обновляем ценник для корзины
// скидка действует только на цены позиций без скидки
func(p *Cart) PriceCalculate() {
	//имеется скидка на всю корзину
	if p.Discount != nil {
		//подсчитываем общую сумму корзины
		positions := p.Positions
		p.Price = 0
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
	} else {
		positions := p.Positions
		//подсчитываем общую сумму корзины с учетом скидки каждой позиции
		p.Price = 0
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
	}
}

// обновление цены у позиции корзины
func(p *Position) PriceCalculate() {
	product := p.Product
	//цена позиции (цена продукта * общее количество)
	p.Price = product.Price * p.Amount

	if p.Discount != nil {
		//если существует скидка у позиции, расчитываем цену со скидкой
		p.Discount.Price = PriceComputer(p.Price, p.Discount, 1)
	}
}
