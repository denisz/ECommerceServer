package cart

import (
	"github.com/dgrijalva/jwt-go"
)

type Item struct {
	ProductID int `json:"productID"`
	Amount int `json:"amount"`
}

type Cart struct {
	Items []Item `json:"items"`
	jwt.StandardClaims
}

type ItemDTO struct {
	ProductID int `json:"productID"`
	Amount int `json:"amount"`
}

type CartDTO struct {
	Items []ItemDTO `json:"items"`
}