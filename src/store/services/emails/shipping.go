package emails

import (
	"github.com/matcornic/hermes"
	"store/models"
	"fmt"
)

// Заказ отправлен
type Shipping struct {
	Order models.Order
}

func (p Shipping) Subject() string {
	return fmt.Sprintf("Заказ № %s", p.Order.Invoice)
}

func (p Shipping) Recipient() string {
	return p.Order.Address.Email
}

func (p Shipping) Email() hermes.Email {
	var table [][]hermes.Entry

	for _, position := range p.Order.Positions {
		table = append(table, []hermes.Entry{
			{Key: "Позиция", Value: fmt.Sprintf("%s x %d", position.Product.Name, position.Amount )},
		})
	}

	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      p.Order.Address.Name,
			Signature: "С уважением",
			Intros: []string{
				"Информируем вас о том, что ваш заказ был отправлен.",
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
					Instructions: "Отслеживать информацию о вашей посылке и получить трек-код вы можете на",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Отследить посылку",
						Link:  "http://95.213.236.60",
					},
				},
			},
			Outros: []string{
				"Будем рады видеть вас в числе постоянных клиентов нашего магазина.",
			},
		},
	}

}
