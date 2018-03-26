package models

// Адрес
type Address struct {
	//Имя получателя
	Name string `json:"name"`

	//Электронная почта
	Email string `json:"email"`

	//Название компании (если указывается рабочий адрес)
	Company string `json:"company"`

	//Улица, номер дома, корпус или строение
	Address string `json:"address"`

	//Район (редко, встречается в английских и ирландских адресах)
	Region string `json:"region"`

	//Город
	City string `json:"city"`

	//Телефон
	Phone string `json:"phone"`

	//Страна
	Country string `json:"country"`

	//Почтовый индекс
	PostalCode string `json:"postalCode"`
}
