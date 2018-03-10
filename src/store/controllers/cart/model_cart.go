package cart

import (
	"github.com/dgrijalva/jwt-go"
	"store/controllers/catalog"
)

type Position struct {
	Amount int `json:"amount"`
	ProductID int `json:"productID"`
}

type Session struct {
	jwt.StandardClaims
	Positions []Position
}

type Cart struct {
	TotalPrice int `json:"totalPrice"`
	Positions []Position `json:"positions"`
	Products []catalog.Product `json:"products"`
}

type Operation string

const (
	OperationInsert Operation = "insert" //Добавить товар
	OperationUpdate Operation = "update" //Обновить товар
	OperationDelete Operation = "delete" //Удалить товар
)

type UpdateDTO struct {
	Amount int `json:"amount"`
	ProductID int `json:"productID"`
	Operation Operation `json:"operation"`
}