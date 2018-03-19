package account

import "time"

type (
	// Пользователь
	User struct {
		ID          int       `storm:"id,increment" json:"id"`
		Email       string    `storm:"index" json:"-"`
		Password    string    `storm:"index" json:"-"`
		Group       string    `storm:"index" json:"-"`
		CreatedAt   time.Time `json:"createdAt"`
		AddressBook []Address `json:"addressBook"`
	}

	// Адрес
	Address struct {
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

	LoginDTO struct {
		Email    string `json:"email" valid:"email"`
		Password string `json:"password"`
	}

	EmailRegisterDTO struct {
		Email string `json:"email" valid:"email"`
	}

	ForgetDTO struct {
		Email string `json:"email" valid:"email"`
	}

	/**
		Алгоритм регистрации нового пользователя

	1. Пользователь вводит Email
	2. Если данный адрес занят пользователь получает ошибку и конфликте username
	3. На указзыный Email уходит письмо с допуском на сайт
	4. Человек регистрируется
	 */
)
