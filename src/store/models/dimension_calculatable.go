package models

//список контейнеров
//расчет размеров корзины
func (p *Cart) DimensionCalculate() Dimension {
	return Dimension{
		Width:  10,
		Height: 10,
		Length: 10,
	}
}

//расчет размеров корзины
func (p *Order) DimensionCalculate() Dimension {
	return Dimension{
		Width:  10,
		Height: 10,
		Length: 10,
	}
}

//расчет объема
func (p *Dimension) VolumeCalculate() int {
	return p.Width * p.Length * p.Height
}
