package models

import (
	"github.com/dgrijalva/jwt-go"
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
	SessionPosition struct {
		// Количество
		Amount int `json:"amount"`

		// Индентификатор
		ProductSKU string `json:"productSKU"`
	}

	// Модель хранения корзины
	Session struct {
		// Позиции
		Positions []SessionPosition

		jwt.StandardClaims
	}

	// Позиция
	Position struct {
		// Цена
		Price int `json:"price"`

		// Скидка
		Discount *Discount `json:"discount"`

		// Количество
		Amount int `json:"amount"`

		// Индентификатор
		ProductSKU string `json:"productSKU"`

		// Описание продукта
		Product Product `json:"product"`
	}

	// Корзина
	Cart struct {
		// Общая цена
		TotalPrice int `json:"totalPrice"`

		// Позиции
		Positions []Position `json:"positions"`
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
