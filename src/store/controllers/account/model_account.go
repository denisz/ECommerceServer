package account

import "time"

type User struct {
	ID int `storm:"id,increment" json:"id"`
	Email string `storm:"index" json:"-"`
	Password string `storm:"index" json:"-"`
	Group string `storm:"index" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	AddressBook []Address `json:"addressBook"`
}

type Address struct {
	Name string `json:"name"` //Имя получателя
	Email string `json:"email"`//Электронная почта
	Company string `json:"company"` //Название компании (если указывается рабочий адрес)
	Address string `json:"address"` //Улица, номер дома, корпус или строение
	Region string `json:"region"` //Район (редко, встречается в английских и ирландских адресах)
	City string `json:"city"` //Город
	Phone string `json:"phone"` //Телефон
	Country string `json:"country"` //Страна
	PostalCode string `json:"postalCode"` //Почтовый индекс
}

type LoginDTO struct {
	Email string `json:"email" valid:"email"`
	Password string `json:"password"`
}

type EmailRegisterDTO struct {
	Email string `json:"email" valid:"email"`
}

type ForgetDTO struct {
	Email string `json:"email" valid:"email"`
}

/**
	Алгоритм регистрации нового пользователя

1. Пользователь вводит Email
2. Если данный адрес занят пользователь получает ошибку и конфликте username
3. На указзыный Email уходит письмо с допуском на сайт
4. Человек регистрируется
 */