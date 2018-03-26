package models

import "time"

type ReportType int

const (
	//Корректирующий отчет
	ReportTypeFix ReportType = 0

	//Расходная накладная
	ReportTypeDelivery ReportType = 1

	//Приходная накладная
	ReportTypeReceiving ReportType = 2
)
type (
	PositionReport struct {
		// Количество
		Amount int `json:"amount"`

		// Индентификатор
		ProductSKU string `json:"productSKU"`
	}

	//Накладная
	AccountReport struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		//Тип
		Type ReportType `storm:"index" json:"type"`

		//Позиции
		Positions []PositionReport `json:"positions"`

		//Оператор
		Operator int `storm:"index" json:"operator"`

		// Время создания
		CreatedAt time.Time `json:"createdAt"`
	}

	//Ведомость наличия на складе
	StockStatusReport struct {
		// Количество
		Amount int `json:"amount"`

		// Индентификатор
		ProductSKU string `storm:"id" json:"productSKU"`

		// Время создания
		UpdatedAt time.Time `json:"updatedAt"`
	}
)