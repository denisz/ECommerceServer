package cdek

import (
	"testing"
	"time"
	"encoding/xml"
	"fmt"
)

/**
<DeliveryRequest account="123" currency="EUR" date="2015-09-29T00:00:00+06:00" number="shop_20150929_0000000023" orderCount="1" secure="123">
  <order DateInvoice="2015-09-29" Number="ORD123023" phone="2255665" recCityCode="270" recipientEmail="" recipientName="Rublijow Ilja" sellerAddress="Ventura Drive 4726,,CA,Salinas" sellerName="Arnold Debbie" sendCityCode="10812" ShipperAddress="Ventura Drive 4726,,CA,Salinas" ShipperName="Arnold Debbie" tariffTypeCode="8">
    <address flat="-" house="-" street="Kirova 94,,Krasnodarskij kraj,GELENDZHIK" />
    <package barCode="AA000000031" number="AA000000031" weight="2500">
      <item amount="4" comment="Harry Poterr tom 1,Fabric" commentEx="Harry Poterr tom 1" cost="1615.55" costEx="24.99" link="-" payment="0.0" paymentEx="0.0" wareKey="ORD123023-1" weight="400" weightBrutto="400" />
    </package>
  </order>
</DeliveryRequest>
 */
func TestSDK_CreateBacklog(t *testing.T) {
	r := DeliveryRequest{
		ForeignDelivery: 1,
		Number:          "shop_20150929_0000000023",
		Date:            time.Now().Format("2006-01-02T15:04:05-07:00"),
		Currency:        CurrencyRUB,
		Account:         accountDev,
		Secure:          CreateSecureSignature(passwordDev, time.Now().Format("2006-01-02T15:04:05-07:00")),
		OrderCount:      1,
		Order: []OrderCreateRequest{
			{
				Number:           "ORD123023",
				DateInvoice:      time.Now().Format("2006-01-02"),
				SendCityCode:     10812,
				RecCityCode:      270,
				SendCityPostCode: "",
				RecCityPostCode:  "",
				Phone:            "2255665",
				RecipientName:    "Rublijow Ilja",
				RecipientEmail:   "",
				SellerAddress:    "Ventura Drive 4726,,CA,Salinas",
				SellerName:       "Arnold Debbie",
				ShipperAddress:   "Ventura Drive 4726,,CA,Salinas",
				ShipperName:      "Arnold Debbie",
				TariffTypeCode:   8,
				ItemsCurrency: CurrencyRUB,
				RecipientCurrency: CurrencyRUB,
				Address: AddressRequest{
					Flat:   "-",
					House:  "-",
					Street: "Kirova 94,,Krasnodarskij kraj,GELENDZHIK",
				},
				Package: PackageRequest{
					Number:  "AA000000031",
					Weight:  2500,
					BarCode: "AA000000031",
					Items: []ItemRequest{
						{
							Amount:       4,
							Comment:      "Harry Potter tom 1,Fabric",
							CommentEx:    "Harry Potter tom 1,Fabric",
							Cost:         1625.55,
							CostEx:       24.99,
							Link:         "",
							Payment:      0.0,
							PaymentEx:    0.0,
							WareKey:      "ORD123023-1",
							Weight:       400,
							WeightBrutto: 400,
						},
					},
				},
			},
		},
	}

	encode, _ := xml.Marshal(r)
	xmlStr := fmt.Sprintf("%s%s", xml.Header, string(encode))

	fmt.Printf("xml: %s \n", xmlStr)

	orderID, err := DebugSDK.CreateBacklog(r)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("New order with id: %d \n", orderID)
}

/**
<?xml version="1.0" encoding="UTF-8" ?>
<DeleteRequest Number="236" Date="2010-10-14" Account=”abc123” Secure=”abcd1234” OrderCount=”2”>
    <Order Number="5403" />
    <Order Number="5404" />
</DeleteRequest>
 */

func TestSDK_DeleteBacklog(t *testing.T) {
	r := DeleteRequest{
		Number:     "236",
		Date:       time.Now().Format("2006-01-02"),
		Account:    accountDev,
		Secure:     CreateSecureSignature(passwordDev, time.Now().Format("2006-01-02")),
		OrderCount: 2,
		Orders: []OrderDeleteRequest{
			{Number: "5403"},
			{Number: "5404"},
		},
	}

	encode, _ := xml.Marshal(r)
	xmlStr := fmt.Sprintf("%s%s", xml.Header, string(encode))

	fmt.Printf("xml: %s", xmlStr)
}
