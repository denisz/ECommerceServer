package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	. "store/models"
)

var CartSecretKey = []byte("sfvDUPC0Cj")
const CartCookieName = "CART_ID"

func ReadSessionFromRequest(c *gin.Context) *Session {
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

func WriteSessionToResponse(c *gin.Context, session *Session) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), session)
	tokenString, err := token.SignedString(CartSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(CartCookieName, tokenString, 7 * 24 * 3600, "/","",false, true)
}

func appendIfNeeded(positions []Position, productSKU string) []Position {
	var exists bool
	for _, v := range positions {
		if exists == false {
			exists = v.ProductSKU == productSKU
		}
	}

	if exists == false {
		positions = append(positions, Position{
			ProductSKU: productSKU,
			Amount: 0,
		})
	}

	return positions
}