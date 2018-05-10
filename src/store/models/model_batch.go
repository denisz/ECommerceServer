package models

import (
	"store/delivery/russiaPost"
	"time"
)

type (
	Batch struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`
		// Повставщик даоставки
		Provider DeliveryProvider `json:"provider"`
		// Сущность сервиса (Почта Росссии)
		PayloadRussiaPost russiaPost.Batch `json:"payloadRussiaPost"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
	}

	PageBatches struct {
		Content []Batch `json:"content"`
		// Курсор
		Cursor
	}
)
