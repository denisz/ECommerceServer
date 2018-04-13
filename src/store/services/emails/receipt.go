package emails

import (
	"github.com/matcornic/hermes"
	"store/models"
	"fmt"
)

// Заказ создан
type Receipt struct {
	Order models.Order
}

func (p Receipt) Subject() string {
	return "Спасибо за ваш заказ"
}

func (p Receipt) Recipient() string {
	return p.Order.Address.Email
}

func (p Receipt) Email() hermes.Email {
	var table [][]hermes.Entry

	for _, position := range p.Order.Positions {
		table = append(table, []hermes.Entry{
			{Key: "Позиция", Value: fmt.Sprintf("%s x %d", position.Product.Name, position.Amount )},
			{Key: "Цена", Value: fmt.Sprintf("%d руб.", position.Total / 100)},
		})
	}

	table = append(table, []hermes.Entry{
		{Key: "Итого", Value: "Цена товара:"},
		{Key: "Цена", Value: fmt.Sprintf("%d руб.", p.Order.ProductPrice / 100)},
	})

	table = append(table, []hermes.Entry{
		{Key: "Итого", Value: "Цена доставки:"},
		{Key: "Цена", Value: fmt.Sprintf("%d руб.", p.Order.DeliveryPrice / 100)},
	})

	table = append(table, []hermes.Entry{
		{Key: "Итого", Value: "Итого к оплате:"},
		{Key: "Цена", Value: fmt.Sprintf("%d руб.", p.Order.Total / 100)},
	})

	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      p.Order.Address.Name,
			Signature: "С уважением",
			Intros: []string{
				"Благодарим Вас за заказ в интернет-магазине Dark Waters",
			},
			Dictionary: []hermes.Entry{
				{Key: "Номер заказа", Value: p.Order.Invoice},
				{Key: "Статус заказа", Value: p.Order.Status.Format()},
				{Key: "Доставка", Value: p.Order.Delivery.Format()},
				{Key: "Адрес", Value: p.Order.Address.Format()},
			},
			Table: hermes.Table{
				Data: table,
				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Цена":  "100px",
					},
					CustomAlignment: map[string]string{
						"Цена":  "left",
						"Итого": "right",
					},
				},
			},
			Actions: []hermes.Action{
				{
					Instructions: "Вы можете проверить статус Вашего заказа:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Заказ",
						Link:  fmt.Sprintf("http://95.213.236.60/order/check/%s", p.Order.Invoice),
					},
				},
			},
			FreeMarkdown: `
**Способы оплаты**

-   кошелек Яндекс Денег № 410014272453392
-   Qiwiкошелек  № 410014272453392
-   Web.Money кошелек  № 410014272453392
-   банковская карта № 4279-3100-1816-4136 (СберБанк)

Если вы оплачиваете заказ через электронный кошелек - обязательно укажите номер вашего заказа в комментарии.  
Если оплачиваете через терминал, кассу банка или банкомат - обязательно сделайте скан/фото чека.

**Платежные реквизиты действительны в течении  суток с момента оформления заказа!**  
При задержке оплаты уточняйте реквизиты на почте [admin@real-pump.com](https://mail.rambler.ru/#/compose/to=admin%2540real-pump.com)

**Подтверждение оплаты**
После совершения оплаты, ответьте на письмо или напишите нам письмо на почту [admin@real-pump.com](https://mail.rambler.ru/#/compose/to=admin%2540real-pump.com), в котором укажите:

- способ оплаты
- скан/фото чека об оплате.
- Номер вашего заказа или фамилию

После подтверждения оплаты в течение нескольких часов мы сформируем вашу посылку и отправим по указанному адресу.  

После отправки заказа мы вышлем трек-код для слежения за посылкой на Ваш email.`,
			Outros: []string{},
		},
	}
}
