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
	// Модель хранения корзины
	Session struct {
		// Корзина
		CardID int

		jwt.StandardClaims
	}

	// Позиция
	Position struct {
		// Цена позиции без скидки
		Price int `json:"price"`

		//скидка
		Discount *Discount `json:"discount"`

		// Количество
		Amount int `json:"amount"`

		// Индентификатор
		ProductSKU string `json:"productSKU"`

		// Описание продукта
		Product *Product `json:"product"`
	}

	// Корзина
	Cart struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Цена корзины без скидок
		Price int `json:"price"`

		// Скидка
		Discount *Discount `json:"discount"`

		// Цена доставки
		DeliveryPrice int `json:"deliveryPrice"`

		// Конечная цена
		TotalPrice int `json:"totalPrice"`

		//Адресс
		Address *Address `json:"address"`

		//Доставка
		Delivery *Delivery `json:"delivery"`

		// Доступные способы доставки
		DeliveryProviders []DeliveryProvider `json:"deliveryProviders"`

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
