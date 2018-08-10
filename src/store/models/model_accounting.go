package models

import "time"

type ReportType string

const (
	// Расходная накладная
	ReportTypeDelivery ReportType = "delivery"

	// Приходная накладная
	ReportTypeReceiving ReportType = "receiving"
)

type SourceReport string

const (
	//Google таблички
	SourceReportGoogle SourceReport = "google"
	//Магазин
	SourceReportStore SourceReport = "store"
	//Ручное создание накладных (возможно потребуется для корекции)
	SourceReportAdmin SourceReport = "admin"
)

type FilterReportWhere string

const (
	//поиск по артикулу
	FilterReportWhereSKU       FilterReportWhere = "sku"
	FilterReportWhereDate      FilterReportWhere = "date"
	FilterReportWhereRangeDate FilterReportWhere = "range_date"
)

type (
	PositionReport struct {
		// Количество
		Amount int `json:"amount"`
		// Индентификатор
		ProductSKU string `json:"productSKU"`
		// Цена (для приходного это цена закупки, для расходного цена продажи)
		Price Price `json:"price"`
	}

	// Накладная
	AccountingReport struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`
		// Тип
		Type ReportType `storm:"index" json:"type"`
		// Позиции
		Positions []PositionReport `json:"positions"`
		// Оператор
		OperatorID int `storm:"index" json:"operatorId"`
		// Время создания
		CreatedAt time.Time `json:"createdAt"`
		// Время последнего обновления
		UpdatedAt time.Time `json:"updatedAt"`
		// Источник накладной (или google табличка или магазин)
		Source SourceReport `json:"source"`
		// Источник инденцификатор
		SourceID string `json:"sourceId"`
	}

	// Фильтр поиска
	FilterAccountingReport struct {
		// Где искать?
		Where FilterReportWhere `json:"where"`
		// Запрос
		Query string `json:"query"`
		// Начальная дата
		StartDate time.Time `json:"start_date"`
		// Конечная дата
		EndDate time.Time `json:"end_date"`
		// Тип
		Type ReportType `json:"type"`
	}

	// Страницы заказов
	PageAccountingReport struct {
		Content []AccountingReport `json:"content"`
		// Курсор
		Cursor
	}

	// Ведомость наличия на складе
	StockStatusReport struct {
		// Количество
		Amount int `json:"amount"`
		// Индентификатор
		ProductSKU string `storm:"id" json:"productSKU"`
		// Время создания
		UpdatedAt time.Time `json:"updatedAt"`
	}
)

/**
	Алгоритм введения складом

Основная сущность - это налкданая. Бывает расходная и приходная.



 */
