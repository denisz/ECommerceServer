package models

import "time"

type (
	// Пользователь
	User struct {
		ID          int       `storm:"id,increment" json:"id"`
		Email       string    `storm:"index" json:"-"`
		Password    string    `storm:"index" json:"-"`
		Group       string    `storm:"index" json:"-"`
		FirstName   string    `json:"firstName"`
		LastName    string    `json:"lastName"`
		CreatedAt   time.Time `json:"createdAt"`
		AddressBook []Address `json:"addressBook"`
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

//добавить админов