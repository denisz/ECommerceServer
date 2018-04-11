package utils

import (
	"github.com/matcornic/hermes"
	"gopkg.in/gomail.v2"
	"store/services/emails/themes"
)

//Cc: (копия, carbon copy) — вторичные получатели письма, которым направляется копия. Они видят и знают о наличии друг друга.
//m.SetAddressHeader("Cc", "d.zaycev@bytexgames.ru", "Denis")
//Bcc: (скрытая копия, blind carbon copy) — скрытые получатели письма, чьи адреса не показываются другим получателям.
//m.SetAddressHeader("Bcc", "d.zaycev@bytexgames.ru", "Denis")


type Email interface {
	Subject() string
	Recipient() string
	Email() hermes.Email
}

var EmailClient = gomail.NewDialer("smtp.yandex.ru", 465, "d.zaycev@bytexgames.ru", "2Q2sminvc")

func CreateBrand() hermes.Hermes {
	return hermes.Hermes{
		// Optional Theme
		Theme: new(themes.Default),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Dark Waters",
			Link: "http://95.213.236.60",
			// Optional product logo
			Logo: "http://95.213.236.60/img/ic_brand.png",
		},
	}
}

func SendEmails(h hermes.Hermes, email Email, recipients []string) error {
	emailBody, err := h.GenerateHTML(email.Email())
	if err != nil {
		return err
	}

	for _, recipient := range recipients {
		m := gomail.NewMessage()
		m.SetHeader("From", "d.zaycev@bytexgames.ru")
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", email.Subject())
		m.SetBody("text/html", emailBody)
		err := EmailClient.DialAndSend(m)
		if err != nil {
			return err
		}
	}

	return nil
}

func SendEmail(h hermes.Hermes, email Email) error {
	emailBody, err := h.GenerateHTML(email.Email())
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "d.zaycev@bytexgames.ru")
	m.SetHeader("To", email.Recipient())
	m.SetHeader("Subject", email.Subject())
	m.SetBody("text/html", emailBody)

	return EmailClient.DialAndSend(m)
}
