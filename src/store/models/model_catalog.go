package models

type DiscountType int32

const (
	// Процентная скидка
	DiscountTypePercentage   DiscountType = 0
	
	// Фиксированная скидка
	DiscountTypeFixedAmount  DiscountType = 1

	// Бесплатная доставка
	DiscountTypeFreeShipping DiscountType = 2
)

type (
	// Скидка
	Discount struct {
		// Тип скидки
		Type DiscountType `json:"type"`

		// Количество
		Amount int `json:"amount"`
	}

	// Категория
	Collection struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Имя категории
		Name string `storm:"index" json:"name"`

		// Изображение
		Picture string `json:"picture"`

		// Артикул
		SKU string `storm:"index" json:"SKU"`
	}

	// Продукт
	Product struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Имя
		Name string `storm:"index" json:"name"`

		// Производитель
		Producer string `storm:"index" json:"producer"`

		// Лекарственная форма (т.е. таблетки, порошок, капли) и тип упаковки
		Factor string `json:"factor"`

		// Форма (20шт по 10мг)
		Form string `json:"form"`

		// Вес
		Weight int `json:"weight"`

		// Артикул
		SKU string `storm:"index" json:"SKU"`

		// Количество
		Quantity int `json:"quantity"`

		// Цена
		Price int `json:"price"`

		// Категория
		CollectionSKU string `storm:"index" json:"collectionSKU"`

		// Список изображений
		Pictures []string `json:"pictures"`

		// Скидка
		Discount *Discount `json:"discount"`
	}

	// Описание
	Notation struct {
		// Основное описание
		Description string `json:"description"`

		// Побочные эффекты
		BadEffect string `json:"badEffect"`
	}

	// Страницы категорий
	PageCollections struct {
		Content []Collection `json:"content"`

		// Курсор
		Cursor
	}

	// Страницы продуктов
	PageProducts struct {
		Content []Product `json:"content"`

		// Курсор
		Cursor
	}
)