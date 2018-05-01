package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestPriceComputer(t *testing.T) {
	var price Price = 1000
	discount := Discount{
		Type: DiscountTypePercentage,
		Amount: 10,
	}
	expected := 900
	actual := PriceComputer(price, &discount, 1)
	assert.Equal(t, expected, actual)
}

func TestInBetween(t *testing.T) {
	assert.True(t, InBetween(2, 1, 3))
	assert.True(t, InBetween(2, 1, math.MaxInt32))
}
