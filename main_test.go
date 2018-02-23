package main

import (
	"testing"
	"github.com/asdine/storm"
	"fmt"
	"math/rand"
	"store/controllers/catalog"
	"gopkg.in/gomail.v2"
	"github.com/matcornic/hermes"
)

func TestMock(t *testing.T) {
	db, err := storm.Open("store.db")
	if err != nil {
		t.Error(err)
		return
	}

	err = db.Drop("store")
	if err != nil {
		t.Error(err)
	}

	c := db.From("store")

	for i := 1; i <= 10; i++ {
		collection := &catalog.Collection{
			Name: fmt.Sprintf("Collection %d", i),
		}

		c.Save(collection)

		for i := 1; i <= 10; i++ {
			product := &catalog.Product{
				Quantity: 10,
				CollectionID: collection.ID,
				Slug: fmt.Sprintf("product_%d", i),
				Name: fmt.Sprintf("Product %d", i),
				Desc: fmt.Sprintf("Description %d", i),
			}

			if rand.Float64() > 0.5 {
				product.Discount = &catalog.Discount{
					Type: catalog.DiscountTypeFixedAmount,
					Amount: 10,
				}
			}

			c.Save(product)
		}
	}

	var collections []catalog.Collection
	err = c.All(&collections)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf(" %v", collections)

	var products []catalog.Product
	err = c.All(&products)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf(" %v", products)
}

func TestSendHermes(t *testing.T) {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					{
						{Key: "Item", Value: "Golang"},
						{Key: "Description", Value: "Open source programming language that makes it easy to build simple, reliable, and efficient software"},
						{Key: "Price", Value: "$10.99"},
					},
					{
						{Key: "Item", Value: "Hermes"},
						{Key: "Description", Value: "Programmatically create beautiful e-mails using Golang."},
						{Key: "Price", Value: "$1.99"},
					},
				},
				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignment: map[string]string{
						"Price": "right",
					},
				},
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
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

	m := gomail.NewMessage()
	m.SetHeader("From", "d.zaycev@bytexgames.ru")
	m.SetHeader("To", "denisxy12@gmail.com")
	//Cc: (копия, carbon copy) — вторичные получатели письма, которым направляется копия. Они видят и знают о наличии друг друга.
	//m.SetAddressHeader("Cc", "d.zaycev@bytexgames.ru", "Denis")
	//Bcc: (скрытая копия, blind carbon copy) — скрытые получатели письма, чьи адреса не показываются другим получателям.
	//m.SetAddressHeader("Bcc", "d.zaycev@bytexgames.ru", "Denis")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", emailBody)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.yandex.ru", 465, "d.zaycev@bytexgames.ru", "2Q2sminvc")

	if err := d.DialAndSend(m); err != nil {
		t.Error(err)
	}

}