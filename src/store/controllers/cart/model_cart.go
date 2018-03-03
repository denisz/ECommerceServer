package cart

import (
	"github.com/dgrijalva/jwt-go"
	"store/controllers/catalog"
)

type Position struct {
	ProductID int `json:"productID"`
	Amount int `json:"amount"`
}

type Session struct {
	jwt.StandardClaims
	Positions []Position
}

type Cart struct {
	Products []catalog.Product `json:"products"`
	Positions []Position `json:"positions"`
}

type ItemDTO struct {
	ProductID int `json:"productID"`
	Amount int `json:"amount"`
}