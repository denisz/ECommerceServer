package pochta

import (
	"testing"
	"fmt"
)

//viktor@otdeldostavok.ru:123456qQ
//9a9mk3FmmR1E84cn7FHMlz9Kjm5NHAC6
//dmlrdG9yQG90ZGVsZG9zdGF2b2sucnU6MTIzNDU2cVE=

/**
	"index-from": "101000",
	"index-to": "644015",
	"mail-category": "ORDINARY",
	"mail-type": "POSTAL_PARCEL",
	"mass": 1000,
	"rcp-pays-shipping": "false",
	"dimension": {
		"height": 2,
		"length": 5,
		"width": 197
	},
	"fragile": "true"
 */

func TestTariff(t *testing.T) {
	token := "9a9mk3FmmR1E84cn7FHMlz9Kjm5NHAC6"
	login := "viktor@otdeldostavok.ru"
	password := "123456qQ"
	pochta := NewPochta(login, password, token)

	r := &DestinationRequest{
		IndexFrom: "430030",
		IndexTo: "141895",
		MailCategory: MailCategorySIMPLE,
		MailType: MailTypeONLINE_PARCEL,
		Dimension: Dimension{
			Width: 10,
			Height: 10,
			Length: 10,
		},
		PaymentMethod:PaymentMethodCASHLESS,
		Mass: 1000,
		Fragile: true,
		WithOrderOfNotice: false,
		WithSimpleNotice: false,
	}

	res, err := pochta.Tariff(r)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v" ,res)
}

