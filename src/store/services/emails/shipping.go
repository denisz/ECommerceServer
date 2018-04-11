package emails

import (
	"github.com/matcornic/hermes"
	"store/models"
)

// Заказ отправлен
type Shipping struct {
	Order models.Order
}

func (p Shipping) Subject() string {
	return "Чек"
}

func (p Shipping) Recipient() string {
	return p.Order.Address.Email
}

func (p Shipping) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      "Игорь",
			Signature: "С уважением",
			Intros: []string{
				"Благодарим Вас за заказ в интернет-магазине Dark Waters",
			},
			Dictionary: []hermes.Entry{
				{Key: "Номер вашего заказа", Value: "1000"},
				{Key: "Статус заказа", Value: "В обработке"},
				{Key: "Способ доставки", Value: "Почта России - Стандарт"},
				{Key: "Адрес доставки", Value: "430030, Россия, Республика Мордовия, Саранск, улица Полежаева, 120 "},
			},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					{
						{Key: "Позиция", Value: "Epistane"},
						{Key: "Количество", Value: "3"},
						{Key: "Артикул", Value: "PG_AB_EPI_90"},
						{Key: "Цена", Value: "10.99 руб."},
					},
					{
						{Key: "Позиция", Value: "Hermes"},
						{Key: "Количество", Value: "3"},
						{Key: "Артикул", Value: "PG_AB_EPI_90"},
						{Key: "Цена", Value: "1.99 руб."},
					},
					{
						{Key: "Позиция", Value: ""},
						{Key: "Количество", Value: ""},
						{Key: "Итого", Value: "Сумма к оплате:"},
						{Key: "Цена", Value: "100 руб."},
					},
				},
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
					},
				},
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Продолжить покупки",
						Link:  "http://95.213.236.60",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

}
