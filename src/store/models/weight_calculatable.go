package models


func (p *Cart) WeightCalculate() int {
	var weight = 0
	for _, position := range p.Positions {
		weight = weight + position.WeightCalculate()
	}

	return weight
}

func (p *Order) WeightCalculate() int {
	var weight = 0
	for _, position := range p.Positions {
		weight = weight + position.WeightCalculate()
	}

	return weight
}

func (p *Position) WeightCalculate() int {
	return p.Product.Weight * p.Amount
}