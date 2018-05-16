package models

import (
	"time"
)

type (
	//Заказ в партии
	BatchOrder struct {
		// ID заказа
		ID int `json:"id"`
		// Цена
		Total Price `json:"price"`
		// Счет
		Invoice string `json:"invoice"`
		//Весовка
		Weight int `json:"weight"`
		// Имя получателя
		RecipientName string `json:"recipientName"`
	}

	// Партия
	Batch struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`
		// Список заказов
		Orders []BatchOrder `json:"orders"`
		// Повставщик даоставки
		Provider DeliveryProvider `json:"provider"`
		// Сущность сервиса (Почта Росссии)
		PayloadRussiaPost []string `json:"payloadRussiaPost"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
	}

	// Фильтр batch
	FilterBatch struct {
		// Где искать?
		Where FilterOrderWhere `json:"where"`
		// Запрос
		Query string `json:"query"`
		// Начальная дата
		StartDate time.Time `json:"start_date"`
		// Конечная дата
		EndDate time.Time `json:"end_date"`
		//провайдер
		Provider DeliveryProvider `json:"provider"`
	}

	PageBatches struct {
		Content []Batch `json:"content"`
		// Курсор
		Cursor
	}
)
