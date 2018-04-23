package store

import (
	"testing"
	"io/ioutil"
	. "store/models"
	"store/utils"
	"store/services/emails"
	"store/services/api"
	"time"
	"github.com/matcornic/hermes"
	"store/services/emails/themes"
)

func TestMainPackage(t *testing.T) {

}

func CreateMockOrder(t *testing.T) Order {
	invoice, err := api.CreateInvoice()
	if err != nil {
		t.Error(err)
	}

	return Order {
		Positions: []Position {
			{
				Amount: 3,
				Total: 3000 * 100,
				Subtotal: 3000 * 100,
				Product: &Product {
					Name: "Test product",
					Producer: "Test Producer",
					Price: 1000 * 100,
				},
			},
			{
				Amount: 3,
				Total: 3000 * 100,
				Subtotal: 3000 * 100,
				Product: &Product {
					Name: "Test product",
					Producer: "Test Producer",
					Price: 1000 * 100,
				},
			},
		},
		ProductPrice: 6000 * 100,
		DeliveryPrice: 300 * 100,
		Total: 6200 * 100,
		Delivery: &Delivery{
			Provider: DeliveryProviderRussiaPost,
			Method: DeliveryMethodRapid,
		},
		Discount: &Discount{
			Type: DiscountTypePercentage,
			Amount: 2.5,
		},
		Status: OrderStatusAwaitingPayment,
		Address: &Address {
			ManualInput: false,
			Address: "город Саранск, улица Полежаева д. 120 кв. 309",
			PostalCode: "430030",
			Phone: "+7 (999) 89-795-61",
			Name: "Zaycev Denis",
			Email: "denisxy12@gmail.com",
			Comment: "Код от домофона 309",
		},
		Invoice: invoice,
		CreatedAt: time.Now(),
		ID: 1,
	}
}

func CreateBrand() hermes.Hermes {
	return hermes.Hermes{
		// Optional Theme
		Theme: new(themes.Default),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Dark Waters",
			Copyright: "Copyright © 2018 Dark Waters. All rights reserved.",
			Link: "http://95.213.236.60",
			// Optional product logo
			Logo: "http://95.213.236.60/img/ic_brand.png",
		},
	}
}

func TestEmailReceipt(t *testing.T) {
	order := CreateMockOrder(t)
	brand := CreateBrand()
	email := emails.Receipt{ Order: order }
	err := utils.SendEmail(brand, email)
	if err != nil {
		t.Error(err)
	}

	emailBody, err := utils.DrawEmail(brand, email)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("/tmp/dat1.html", []byte(emailBody), 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestEmailProcessing(t *testing.T) {
	order := CreateMockOrder(t)
	brand := CreateBrand()
	email := emails.Processing{ Order: order }
	err := utils.SendEmail(brand, email)
	if err != nil {
		t.Error(err)
	}

	emailBody, err := utils.DrawEmail(brand, email)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("/tmp/dat1.html", []byte(emailBody), 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestEmailShipping(t *testing.T) {
	order := CreateMockOrder(t)
	brand := CreateBrand()
	email := emails.Shipping{ Order: order }
	err := utils.SendEmail(brand, email)
	if err != nil {
		t.Error(err)
	}

	emailBody, err := utils.DrawEmail(brand, email)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("/tmp/dat1.html", []byte(emailBody), 0644)
	if err != nil {
		t.Error(err)
	}
}