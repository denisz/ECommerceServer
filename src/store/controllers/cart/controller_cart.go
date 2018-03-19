package cart

import (
	"net/http"
	"store/controllers/common"
	"github.com/gin-gonic/gin"
	"store/controllers/catalog"
	"github.com/cznic/mathutil"
)

type Controller struct {
	common.Controller
}

func (p *Controller) IndexPOST(c *gin.Context) {
	session := readCartFromRequest(c)
	c.JSON(http.StatusOK, &Cart{
		Positions: session.Positions,
	})
}

// Детальная информация корзины
func (p *Controller) DetailPOST(c *gin.Context) {
	session := readCartFromRequest(c)
	cart := &Cart{}

	for _, v := range session.Positions {
		var product catalog.Product
		err := p.GetStoreNode().One("SKU", v.ProductSKU, &product)
		if err != nil {
			continue
		}
		cart.Products = append(cart.Products, product)
		cart.Positions = append(cart.Positions, v)
	}

	c.JSON(http.StatusOK, cart)
}

/// Обновить позицию
func (p *Controller) UpdatePOST(c *gin.Context) {
	var json UpdateDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		var positions []Position
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
				var product catalog.Product
				err := p.GetStoreNode().One("SKU", v.ProductSKU, &product)
				if err != nil {
					continue
				}
				v.Amount = mathutil.Clamp(v.Amount, 0, product.Quantity)
				positions = append(positions, v)
			}
		}

		session.Positions = positions
		writeCartToResponse(c, session)
		c.JSON(http.StatusOK, &Cart{
			Positions: positions,
		})
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
