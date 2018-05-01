package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	. "store/models"
)

var SessionSecretKey = []byte("sfvDUPC0Cj") //TODO: REMOVE
const SessionCookieName = "SESSION_ID"

func ReadSessionFromRequest(c *gin.Context) *SessionClaims {
	session := &SessionClaims{}
	tokenString, err := c.Cookie(SessionCookieName)
	if err != nil {
		return session
	}

	jwt.ParseWithClaims(tokenString, session, func(token *jwt.Token) (interface{}, error) {
		return SessionSecretKey, nil
	})

	return session
}

func WriteSessionToResponse(c *gin.Context, session *SessionClaims) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), session)
	tokenString, err := token.SignedString(SessionSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(SessionCookieName, tokenString, 7 * 24 * 3600, "/","",false, true)
}
