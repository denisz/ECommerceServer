package russiaPost

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
	pochta := NewClient(login, password, token, true)

	r := &DestinationRequest{
		Mass: 2000,
		IndexFrom: "200961",
		IndexTo: "430030",
		MailType: MailTypeONLINE_PARCEL,
		MailCategory: MailCategoryORDINARY,
		PaymentMethod: PaymentMethodCASHLESS,
		Dimension: Dimension{
			Width: 10,
			Height: 10,
			Length: 10,
		},
		Fragile: false,
		DeclareValue: 3000,
		WithSimpleNotice: false,
		WithOrderOfNotice: false,
	}

	res, err := pochta.Tariff(r)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v" ,res)
}

