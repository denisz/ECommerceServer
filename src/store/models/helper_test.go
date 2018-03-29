package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPriceComputer(t *testing.T) {
	price := 1000
	discount := Discount{
		Type: DiscountTypePercentage,
		Amount: 10,
	}
	expected := 900
	actual := PriceComputer(price, &discount, 1)
	assert.Equal(t, expected, actual)
}
