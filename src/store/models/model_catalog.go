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
		Amount float64 `json:"amount"`

		// Цена со скидкой
		Price Price `json:"price"`
	}

	// Категория
	Collection struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Имя категории
		Name string `json:"name"`

		// Изображение
		Picture string `json:"picture"`

		// Артикул
		SKU string `storm:"unique" json:"SKU"`
	}

	// Продукт
	Product struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Имя
		Name string `json:"name"`

		// Производитель
		Producer string `storm:"index" json:"producer"`

		// Лекарственная форма (т.е. таблетки, порошок, капли) и тип упаковки
		Factor string `json:"factor"`

		// Форма (20шт по 10мг)
		Form string `json:"form"`

		// Вес
		Weight int `json:"weight"`

		// Артикул
		SKU string `storm:"unique" json:"SKU"`

		// Количество
		Quantity int `json:"quantity"`

		// Цена
		Price Price `json:"price"`

		// Скидка
		Discount *Discount `json:"discount"`

		//Минимальное количество в корзине
		MinQtyAllowed int `json:"-"`

		//Максимальное количество в корзине
		MaxQtyAllowed int `json:"-"`

		// Категория
		CollectionSKU string `storm:"index" json:"collectionSKU"`

		// Список изображений
		Pictures []string `json:"pictures"`

		// Линейные размеры
		Dimension Dimension `json:"dimension"`
	}

	// Описание
	Notation struct {
		// Индентификатор
		ID int `storm:"id,increment" json:"id"`

		// Артикул
		SKU string `storm:"index" json:"SKU"`

		// Основное описание
		Description string `json:"description"`

		// Состав
		Composition string `json:"composition"`

		// Исследования
		Research string `json:"research"`

		// Рекомендации
		Prescribing string `json:"prescribing"`

		// Эффекты
		Effects string `json:"effects"`

		// Матрица
		Matrix string `json:"matrix"`
	}

	// Фильтр поиска
	FilterCatalog struct {
		Query string `json:"query"`
		CollectionSKU string `json:"collectionSKU"`
		Producer string `json:"producer"`
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
