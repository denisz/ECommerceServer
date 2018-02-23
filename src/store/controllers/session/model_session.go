package session

import (
	"github.com/dgrijalva/jwt-go"
	"store/controllers/account"
)

type Session struct {
	UserID  int
	Address account.Address
	jwt.StandardClaims
}

