package cart

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)


var SecretKey = []byte("sfvDUPC0Cj")
const CookieName = "CART_ID"


func readCartFromRequest(c *gin.Context) *Session {
	session := &Session{}
	tokenString, err := c.Cookie(CookieName)
	if err != nil {
		return session
	}

	jwt.ParseWithClaims(tokenString, session, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	return session
}

func writeCartToResponse(c *gin.Context, session *Session) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), session)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(CookieName, tokenString, 7 * 24 * 3600, "/","",false, true)
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