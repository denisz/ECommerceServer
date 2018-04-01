package models

// Адрес
type Address struct {
	//Имя получателя
	FirstName string `json:"firstName"`

	//Фамилия получателя
	LastName string `json:"lastName"`

	//Отчество получателя
	MiddleName string `json:"middleName"`

	//Электронная почта
	Email string `json:"email"`

	//Телефон
	Phone string `json:"phone"`

	//Улица, номер дома, корпус или строение
	Address string `json:"address"`

	//Ручной ввод
	ManualInput bool `json:"manualInput"`

	//Страна
	Country string `json:"country"`

	//Регион
	Region string `json:"region"`

	//Область, район
	District string `json:"district"`

	//Город/Деревня
	City string `json:"city"`

	//Улица
	Street string `json:"street"`

	//Дом
	House string `json:"house"`

	//Квартира
	Room string `json:"room"`

	//Почтовый индекс
	PostalCode string `json:"postalCode"`

	//Координаты
	GeoPoint []float64 `json:"geopoint"`
}
