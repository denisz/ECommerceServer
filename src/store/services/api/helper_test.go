package api

import (
	"testing"
	"fmt"
	. "store/models"
)

func TestHelper(t *testing.T) {
	price := GetPriceWithDiscount(2500, &Discount {
			Type: DiscountTypePercentage,
			Amount: 15,
	}, 1)

	fmt.Printf("%v", price)
}
