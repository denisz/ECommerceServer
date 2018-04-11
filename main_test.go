package main

import (
	"testing"
	"github.com/matcornic/hermes"
	"io/ioutil"
	"store/services/emails/themes"
)

func TestSendHermes(t *testing.T) {
	// Configure hermes by setting a theme and your product info
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		Theme: new(themes.Default),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Dark Waters",
			Link: "http://95.213.236.60",
			// Optional product logo
			Logo: "http://95.213.236.60/img/ic_brand.png",
			Copyright: "Copyright",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Greeting: "Здравствуйте",
			Name: "Игорь",
			Signature: "С уважением",
			Intros: []string{
				"Благодарим Вас за заказ в интернет-магазине Dark Waters",
			},
			Dictionary: []hermes.Entry{
				{Key: "Номер вашего заказа", Value: "1000"},
				{Key: "Статус заказа", Value: "В обработке" },
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
						"Позиция":  "20%",
						"Цена": "15%",
						"Артикул": "15%",
						"Количество": "15%",
						"Итого": "85%",
					},
					CustomAlignment: map[string]string{
						"Цена": "right",
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

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	err = ioutil.WriteFile("/tmp/dat1", []byte(emailBody), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	//m := gomail.NewMessage()
	//m.SetHeader("From", "d.zaycev@bytexgames.ru")
	//m.SetHeader("To", "denisxy12@gmail.com")
	////Cc: (копия, carbon copy) — вторичные получатели письма, которым направляется копия. Они видят и знают о наличии друг друга.
	////m.SetAddressHeader("Cc", "d.zaycev@bytexgames.ru", "Denis")
	////Bcc: (скрытая копия, blind carbon copy) — скрытые получатели письма, чьи адреса не показываются другим получателям.
	////m.SetAddressHeader("Bcc", "d.zaycev@bytexgames.ru", "Denis")
	//m.SetHeader("Subject", "Hello!")
	//m.SetBody("text/html", emailBody)
	////m.Attach("/home/Alex/lolcat.jpg")
	//
	//d := gomail.NewDialer("smtp.yandex.ru", 465, "d.zaycev@bytexgames.ru", "2Q2sminvc")
	//
	//if err := d.DialAndSend(m); err != nil {
	//	t.Error(err)
	//}
}