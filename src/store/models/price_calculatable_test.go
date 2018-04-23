package models

import "testing"
import "github.com/stretchr/testify/assert"


func TestProductPriceCalculate(t *testing.T) {
	product := Product{
		Discount: &Discount{
			Type: DiscountTypePercentage,
			Amount: 10,
		},
		Price: 1000,
	}

	product.PriceCalculate()
	assert.Equal(t, 900, product.Discount.Price)
	assert.Equal(t, 1000, product.Price)
}

func TestProductFloatPriceCalculate(t *testing.T) {
	product := Product{
		Discount: &Discount{
			Type: DiscountTypePercentage,
			Amount: 2.5,
		},
		Price: 1000,
	}

	product.PriceCalculate()
	assert.Equal(t, 975, product.Discount.Price)
	assert.Equal(t, 1000, product.Price)
}

func TestCartPriceCalculate(t *testing.T) {
	product := Product{
		Price: 1000,
		Discount: &Discount {
			Type: DiscountTypePercentage,
			Amount: 10,
		},
	}

	cart := Cart{
		Discount:&Discount{
			Type: DiscountTypePercentage,
			Amount: 10,
		},
		Positions: []Position{
			{
				Amount: 2,
				Discount: &Discount {
					Type: DiscountTypePercentage,
					Amount: 10,
				},
				ProductSKU: "test",
				Product: &product,
			},
			{
				Amount: 3,
				ProductSKU: "test2",
				Discount: &Discount {
					Type: DiscountTypePercentage,
					Amount: 10,
				},
				Product: &product,
			},
		},
	}

	cart.PriceCalculate()
	assert.Equal(t, 2000, cart.Positions[0].Subtotal)
	assert.Equal(t, 1800, cart.Positions[0].Discount.Price)
	assert.Equal(t, 3000, cart.Positions[1].Subtotal)
	assert.Equal(t, 2700, cart.Positions[1].Discount.Price)
	assert.Equal(t, 5000, cart.Total)
	assert.Equal(t, 4500, cart.Discount.Price)
}

func TestPositionPriceCalculate(t *testing.T) {
	product := Product{
		Price: 1000,
		Discount: &Discount {
			Type: DiscountTypePercentage,
			Amount: 10,
		},
	}

	position := Position{
		Amount: 2,
		ProductSKU: "test",
		Product: &product,
		Discount: &Discount {
			Type: DiscountTypePercentage,
			Amount: 10,
		},
	}
	position.PriceCalculate()
	assert.Equal(t, 1800, position.Subtotal)
	assert.Equal(t, 1620, position.Discount.Price)
}

func TestPositionPriceCalculateWithoutDiscount(t *testing.T) {
	product := Product{
		Price: 1000,
	}

	position := Position{
		Amount: 2,
		ProductSKU: "test",
		Product: &product,
	}
	position.PriceCalculate()
	assert.Equal(t, 2000, position.Subtotal)
}
