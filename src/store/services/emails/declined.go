package emails

import (
"github.com/matcornic/hermes"
"store/models"
	"fmt"
)

// Заблоикрован заказ
type Declined struct {
	Order models.Order
}

func (p Declined) Subject() string {
	return fmt.Sprintf("Заказ № %s", p.Order.Invoice)
}

func (p Declined) Recipient() string {
	return p.Order.Address.Email
}

func (p Declined) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Greeting:  "Здравствуйте",
			Name:      p.Order.Address.Name,
			Signature: "С уважением",
			Intros: []string{
				fmt.Sprintf("Информируем вас о том, что ваш заказ %s был отменен. ", p.Order.Invoice),
			},
		},
	}

}


