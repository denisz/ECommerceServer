package cart

import (
	"github.com/dgrijalva/jwt-go"
	"store/controllers/catalog"
)

type Operation string

const (
	// Добавить товар
	OperationInsert Operation = "insert"

	// Обновить товар
	OperationUpdate Operation = "update"

	// Удалить товар
	OperationDelete Operation = "delete"
)

type (
	// Позиция
	Position struct {
		// Количество
		Amount int `json:"amount"`

		// Индентификатор
		ProductSKU string `json:"productSKU"`
	}

	// Модель хранения корзины
	Session struct {
		// Позиции
		Positions []Position

		jwt.StandardClaims
	}

	// Корзина
	Cart struct {
		// Общая цена
		TotalPrice int `json:"totalPrice"`

		// Позиции
		Positions []Position `json:"positions"`

		// Описание продуктов
		Products []catalog.Product `json:"products"`
	}

	// Модель обмена данными
	UpdateDTO struct {
		// Количество
		Amount int `json:"amount"`

		// Продукт
		ProductSKU string `json:"productSKU"`

		// Операци
		Operation Operation `json:"operation"`
	}
)
