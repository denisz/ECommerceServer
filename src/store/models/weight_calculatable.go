package models


func (p *Cart) WeightCalculate() Weight {
	var weight = 0
	for _, position := range p.Positions {
		weight = weight + (position.Product.Weight * position.Amount)
	}

	return Weight(weight)
}

func (p *Order) WeightCalculate() Weight {
	var weight = 0
	for _, position := range p.Positions {
		weight = weight + (position.Product.Weight * position.Amount)
	}

	return Weight(weight)
}

func (p *Position) WeightCalculate() Weight {
	return Weight(p.Product.Weight * p.Amount)
}