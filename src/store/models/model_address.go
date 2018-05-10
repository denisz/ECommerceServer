package models

// Адрес
type Address struct {
	// Имя получателя
	Name string `json:"name"`
	// Электронная почта
	Email string `json:"email"`
	// Телефон
	Phone string `json:"phone"`
	// Улица, номер дома, корпус или строение
	Address string `json:"address"`
	// Ручной ввод
	ManualInput bool `json:"manualInput"`
	// Страна
	Country string `json:"country"`
	// Регион
	Region string `json:"region"`
	// Область, район
	District string `json:"district"`
	// Город/Деревня
	City string `json:"city"`
	// Улица
	Street string `json:"street"`
	// Дом
	House string `json:"house"`
	// Корпус
	Building string `json:"building"`
	// Квартира
	Room string `json:"room"`
	// Комментарий
	Comment string `json:"comment"`
	// Почтовый индекс
	PostalCode string `json:"postalCode"`
	// Праавильный ли указан индекс
	UserInvalidIndex bool `json:"userInvalidIndex"`
	// Координаты
	GeoPoint []float64 `json:"geopoint"`
}
