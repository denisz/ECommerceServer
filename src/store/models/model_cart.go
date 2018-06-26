package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Operation string

type DeliveryMethods map[DeliveryProvider][]DeliveryMethod

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
	SessionClaims struct {
		// Корзина
		CartID int
		// JWT
		jwt.StandardClaims
	}

	// Позиция
	Position struct {
		// Цена позиции без скидки
		Subtotal Price `json:"subtotal"`
		// Окончательная цена
		Total Price `json:"total"`
		// Скидка
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
		Subtotal Price `json:"subtotal"`
		// Скидка
		Discount *Discount `json:"discount"`
		//Цена товаров
		ProductPrice Price `json:"productPrice"`
		// Цена доставки
		DeliveryPrice Price `json:"deliveryPrice"`
		// Минимальное время доставки в днях
		DeliveryPeriodMin int `json:"deliveryPeriodMin "`
		// Максимальное время доставки в днях
		DeliveryPeriodMax int `json:"deliveryPeriodMax "`
		// Окончательная цена
		Total Price `json:"total"`
		// Адрес
		Address *Address `json:"address"`
		// Доставка
		Delivery *Delivery `json:"delivery"`
		// Доступные провайдеры доставки
		DeliveryProviders []DeliveryProvider `json:"deliveryProviders"`
		// Доступные методы доставки
		DeliveryMethods DeliveryMethods `json:"deliveryMethods"`
		// Позиции
		Positions []Position `json:"positions"`
		// Последний заказ
		Invoice string `json:"invoice"`
	}

	// Модель обмена данными
	CartUpdateRequest struct {
		// Количество
		Amount int `json:"amount"`
		// Продукт
		ProductSKU string `json:"productSKU"`
		// Операци
		Operation Operation `json:"operation"`
	}
)
