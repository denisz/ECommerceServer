package models

type Weight int

func (p Weight) Gram() int {
	return int(p)
}

func (p Weight) Kilos() float64 {
	return float64(p) / float64(1000)
}
