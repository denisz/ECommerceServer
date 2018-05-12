package models

import (
	"time"
)

type (
	Batch struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`
		//список заказов
		IDs []int
		// Повставщик даоставки
		Provider DeliveryProvider `json:"provider"`
		// Сущность сервиса (Почта Росссии)
		PayloadRussiaPost []int `json:"payloadRussiaPost"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
	}

	PageBatches struct {
		Content []Batch `json:"content"`
		// Курсор
		Cursor
	}
)
