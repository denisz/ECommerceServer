package emails

import (
"github.com/matcornic/hermes"
"store/models"
)

// Заблоикрован заказ
type Ban struct {
	Order models.Order
}

func (p Ban) Subject() string {
	return "Информация о заказе"
}

func (p Ban) Recipient() string {
	return p.Order.Address.Email
}

func (p Ban) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      p.Order.Address.Name,
			Signature: "С уважением",
			FreeMarkdown: `
# Услуга не доступна!

Сожалеем. но оформление Вами этого заказа невозможно, в течении следующих 24 часов, поскольку за последние сутки Вами было сделано два неоплаченных заказа.
Предлагаем либо оплатить один из ранее оформленных заказов, либо  сделать новый заказ через 24 часа.

Извините за неудобство.
			`,
		},
	}

}

