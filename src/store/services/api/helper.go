package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	. "store/models"
	"math"
)


var CartSecretKey = []byte("sfvDUPC0Cj")
const CartCookieName = "CART_ID"


func readCartFromRequest(c *gin.Context) *Session {
	session := &Session{}
	tokenString, err := c.Cookie(CartCookieName)
	if err != nil {
		return session
	}

	jwt.ParseWithClaims(tokenString, session, func(token *jwt.Token) (interface{}, error) {
		return CartSecretKey, nil
	})

	return session
}

func writeCartToResponse(c *gin.Context, session *Session) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), session)
	tokenString, err := token.SignedString(CartSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(CartCookieName, tokenString, 7 * 24 * 3600, "/","",false, true)
}

func appendIfNeeded(positions []SessionPosition, productSKU string) []SessionPosition {
	var exists bool
	for _, v := range positions {
		if exists == false {
			exists = v.ProductSKU == productSKU
		}
	}

	if exists == false {
		positions = append(positions, SessionPosition{
			ProductSKU: productSKU,
			Amount: 0,
		})
	}

	return positions
}


func GetPriceWithDiscount(price int, discount *Discount, amount int) int {
	if discount != nil {
		switch discount.Type {
		case DiscountTypePercentage:
			sale := float64(price * discount.Amount / 100)
			return (price -  int(math.Floor(sale))) * amount
		case DiscountTypeFixedAmount:
			return (price - discount.Amount) * amount
		}
	}

	return price * amount
}

