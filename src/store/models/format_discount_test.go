package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDiscountFormat(t *testing.T) {
	discount := &Discount{
		Amount: 2.5,
		Type: DiscountTypePercentage,
	}

	assert.Equal(t, discount.ToFormat(), "2.5%")

	discount = &Discount{
		Amount: 2.5,
		Type: DiscountTypeFixedAmount,
	}

	assert.Equal(t, discount.ToFormat(), "2.5 руб.")
}