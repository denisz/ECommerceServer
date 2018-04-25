package emails

import (
"github.com/matcornic/hermes"
)

// Заблоикрован заказ
type Ban struct {
	EmailRecipient string
	NameRecipient string
}

func (p Ban) Subject() string {
	return "Информация о заказе"
}

func (p Ban) Recipient() string {
	return p.EmailRecipient
}

func (p Ban) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      p.NameRecipient,
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

