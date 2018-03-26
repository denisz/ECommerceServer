package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cznic/mathutil"
	. "store/models"
)

type ControllerCart struct {
	Controller
}

func (p *ControllerCart) IndexPOST(c *gin.Context) {
	session := readCartFromRequest(c)
	c.JSON(http.StatusOK, gin.H{
		"Positions": session.Positions,
	})
}

func (p *ControllerCart) GetDetailCart(session *Session) *Cart {
	cart := Cart{}
	for _, v := range session.Positions {
		if len(v.ProductSKU) == 0 || v.Amount <= 0 {
			continue
		}

		var product Product
		err := p.GetCatalog().One("SKU", v.ProductSKU, &product)
		if err != nil {
			continue
		}

		position := Position{
			Product: product,
			Amount: v.Amount,
			Price: product.Price,
			ProductSKU: v.ProductSKU,
			Discount: product.Discount,
		}

		price := GetPriceWithDiscount(product.Price, product.Discount, v.Amount)
		cart.Positions = append(cart.Positions, position)
		cart.TotalPrice = cart.TotalPrice + price
	}

	return &cart
}

// Детальная информация корзины
func (p *ControllerCart) DetailPOST(c *gin.Context) {
	session := readCartFromRequest(c)
	c.JSON(http.StatusOK, p.GetDetailCart(session))
}

/// Обновить позицию
func (p *ControllerCart) UpdatePOST(c *gin.Context) {
	var json UpdateDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		var positions []SessionPosition
		session := readCartFromRequest(c)
		origPositions := appendIfNeeded(session.Positions, json.ProductSKU)

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
		writeCartToResponse(c, session)
		c.JSON(http.StatusOK, p.GetDetailCart(session))
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
