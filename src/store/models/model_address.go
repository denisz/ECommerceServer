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

	//Страна
	Country string `json:"country"`

	//Район
	District string `json:"district"`

	//Город/Деревня
	City string `json:"city"`

	//Почтовый индекс
	PostalCode string `json:"postalCode"`

	//Координаты
	LatLon []float64 `json:"$latlon"`
}
