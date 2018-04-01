package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cznic/mathutil"
	. "store/models"
	"math"
)

type ControllerCart struct {
	Controller
}

func (p *ControllerCart) IndexPOST(c *gin.Context) {
	session := ReadCartFromRequest(c)
	c.JSON(http.StatusOK, gin.H{
		"Positions": session.Positions,
	})
}

/**
	Скидки на корзину:
	С 6 до 10 тыс 2%
	С 10 до 20 тыс 5%
	Свыше 20 тыс. 7%
 */
func (p *ControllerCart) GetDetailCart(session *Session) *Cart {
	cart := Cart{
		Address: session.Address,
	}
	for _, v := range session.Positions {
		if len(v.ProductSKU) == 0 || v.Amount <= 0 {
			continue
		}

		var product Product
		err := p.GetCatalog().One("SKU", v.ProductSKU, &product)
		if err != nil {
			continue
		}

		position := Position {
			Product: product,
			Amount: v.Amount,
			ProductSKU: v.ProductSKU,
			Discount: product.Discount,
		}

		cart.Positions = append(cart.Positions, position)
	}

	cart.PriceCalculate()

	if InBetween(cart.Price, 6000, 10000)  {
		cart.Discount = &Discount {
			Type: DiscountTypePercentage,
			Amount: 2,
		}
		cart.PriceCalculate()
	}

	if InBetween(cart.Price, 10000, 20000)  {
		cart.Discount = &Discount {
			Type: DiscountTypePercentage,
			Amount: 5,
		}
		cart.PriceCalculate()
	}

	if InBetween(cart.Price, 20000, math.MaxInt32)  {
		cart.Discount = &Discount {
			Type: DiscountTypePercentage,
			Amount: 7,
		}
		cart.PriceCalculate()
	}

	return &cart
}

// Детальная информация корзины
func (p *ControllerCart) DetailPOST(c *gin.Context) {
	session := ReadCartFromRequest(c)
	c.JSON(http.StatusOK, p.GetDetailCart(session))
}

/// Обновить позицию
func (p *ControllerCart) UpdatePOST(c *gin.Context) {
	var json UpdateDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		var positions []SessionPosition
		session := ReadCartFromRequest(c)
		origPositions := AppendIfNeeded(session.Positions, json.ProductSKU)

		for _, v := range origPositions {
			if v.ProductSKU == json.ProductSKU {
				switch json.Operation {
				case OperationInsert:
					v.Amount = v.Amount + json.Amount
				case OperationUpdate:
					v.Amount = json.Amount
				case OperationDelete:
					v.Amount = 0
				}
			}

			if v.Amount > 0 {
				if len(v.ProductSKU) == 0 {
					continue
				}

				var product Product
				err := p.GetCatalog().One("SKU", v.ProductSKU, &product)
				if err != nil {
					continue
				}
				v.Amount = mathutil.Clamp(v.Amount, 0, product.Quantity)
				positions = append(positions, v)
			}
		}

		session.Positions = positions
		WriteCartToResponse(c, session)
		c.JSON(http.StatusOK, p.GetDetailCart(session))
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//Сохраняем адрес в сессии для будующих покупок
func (p *ControllerCart) UpdateAddressPOST(c *gin.Context) {
	var json Address

	if err := c.ShouldBindJSON(&json); err == nil {
		session := ReadCartFromRequest(c)
		session.Address = &json
		WriteCartToResponse(c, session)
		c.JSON(http.StatusOK, p.GetDetailCart(session))
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
