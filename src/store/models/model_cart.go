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

		// Адрес
		Address *Address

		jwt.StandardClaims
	}

	// Позиция
	Position struct {
		// Цена позиции без скидки
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
		// Цена корзины без скидок
		Price int `json:"price"`

		//Адресс
		Address *Address `json:"address"`

		//Скидка
		Discount *Discount `json:"discount"`

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
