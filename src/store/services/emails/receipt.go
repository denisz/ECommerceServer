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
			{Key: "Позиция", Value: position.Product.Name},
			{Key: "Количество", Value: fmt.Sprintf("%d", position.Amount)},
			{Key: "Артикул", Value:  position.Product.SKU},
			{Key: "Цена", Value: fmt.Sprintf("%d руб.", position.Total / 100)},
		})
	}

	table = append(table, []hermes.Entry{
		{Key: "Позиция", Value: ""},
		{Key: "Количество", Value: ""},
		{Key: "Итого", Value: "Цена товара:"},
		{Key: "Цена", Value: fmt.Sprintf("%d руб.", p.Order.Subtotal / 100)},
	})

	table = append(table, []hermes.Entry{
		{Key: "Позиция", Value: ""},
		{Key: "Количество", Value: ""},
		{Key: "Итого", Value: "Цена доставки:"},
		{Key: "Цена", Value: fmt.Sprintf("%d руб.", p.Order.DeliveryPrice / 100)},
	})

	table = append(table, []hermes.Entry{
		{Key: "Позиция", Value: ""},
		{Key: "Количество", Value: ""},
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
				{Key: "Номер вашего заказа", Value: p.Order.Invoice},
				{Key: "Статус заказа", Value: p.Order.Status.Format()},
				{Key: "Способ доставки", Value: p.Order.Delivery.Format()},
				{Key: "Адрес доставки", Value: p.Order.Address.Format()},
			},
			Table: hermes.Table{
				Data: table,
				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Позиция":    "20%",
						"Цена":       "15%",
						"Артикул":    "15%",
						"Количество": "15%",
						"Итого":      "85%",
					},
					CustomAlignment: map[string]string{
						"Цена":  "right",
						"Итого": "right",
						"Количество": "center",
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
			Outros: []string{
				"Тут будет текст с оплатой",
			},
		},
	}
}
