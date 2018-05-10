package models

type Price int

func (p Price) RUB() Price {
	return p / 100
}

func (p Price) Cent() Price {
	return p % 100
}

const MaxPrice Price = 1<<31 - 1
const MinPrice Price = -1 << 31
