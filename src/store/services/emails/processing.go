package emails

import (
	"store/models"
	"github.com/matcornic/hermes"
	"fmt"
)

type Processing struct {
	Order models.Order
}

func (p Processing) Subject() string {
	return "Спасибо за ваш заказ"
}

func (p Processing) Recipient() string {
	return p.Order.Address.Email
}


func (p Processing) Email() hermes.Email {
	var table [][]hermes.Entry

	for _, position := range p.Order.Positions {
		table = append(table, []hermes.Entry{
			{Key: "Позиция", Value: fmt.Sprintf("%s x %d", position.Product.Name, position.Amount )},
			{Key: "Цена", Value: fmt.Sprintf("%d руб.", position.Total / 100)},
		})
	}

	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      p.Order.Address.Name,
			Signature: "С уважением",
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
Ваша оплата получена.  
Спасибо за заказ в Нашем магазине.  
Отправка заказов осуществляется каждый день, кроме воскресенья.  
В течении дня, Вам будет отправлена ссылка для отслеживания посылки. Информация об отслеживании будет доступна на следующий день.  
Большая просьба не писать сообщения с вопросом:  
- Выслали ли мне посылку?  
- Когда мне придет трек код?  
И тому подобные. Спасибо.`,
			Outros: []string{},
		},
	}
}

