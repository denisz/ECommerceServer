package api

import (
	. "store/models"
	"store/delivery/russiaPost"
	"fmt"
)

func AppendIfNeeded(positions []Position, productSKU string) []Position {
	var exists bool
	for _, v := range positions {
		if exists == false {
			exists = v.ProductSKU == productSKU
		}
	}

	if exists == false {
		positions = append(positions, Position{
			ProductSKU: productSKU,
			Amount:     0,
		})
	}

	return positions
}

func NormalizeAddressForRussiaPost(address *Address) (*russiaPost.NormalizeAddress, error) {
	req := russiaPost.NormalizeAddressRequest{
		ID:             "Addr 1",
		OriginalString: address.Normalize(),
	}

	return russiaPost.DefaultClient.NormalizeAddress(req)
}

func CheckValidAddress(address *Address) error {
	res, err := NormalizeAddressForRussiaPost(address)
	if err != nil {
		return err
	}

	err = russiaPost.CheckValidateAddress(res)
	if err != nil {
		return err
	}

	return nil
}

func CreateOrderInToRussiaPost(order *Order) error {
	if order.Delivery.Provider != DeliveryProviderRussiaPost {
		return ErrNotSupportedMethod
	}
	//формируем заказ в поставщике доставки
	normalizeAddress, err := NormalizeAddressForRussiaPost(order.Address)
	if err != nil {
		return err
	}

	mailType := russiaPost.MailTypeONLINE_PARCEL

	switch order.Delivery.Method {
	case DeliveryMethodEMC:
		mailType = russiaPost.MailTypePARCEL_CLASS_1
	case DeliveryMethodRapid:
		mailType = russiaPost.MailTypeBUSINESS_COURIER
	case DeliveryMethodStandard:
		mailType = russiaPost.MailTypeONLINE_PARCEL
	}

	dimension := order.DimensionCalculate()
	request := russiaPost.CreateOrderRequestWithAddress(normalizeAddress)
	request.PostOfficeCode = "430005"
	request.Dimension = russiaPost.Dimension{
		Width:  dimension.Width,
		Height: dimension.Height,
		Length: dimension.Length,
	}
	request.GivenName = order.Address.Name
	request.OrderNum = order.Invoice
	//request.TelAddress = order.Address.Phone
	request.RecipientName = order.Address.Name
	request.BrandName = "DarkWaters"
	request.MailType = mailType
	request.MailCategory = russiaPost.MailCategoryORDINARY
	request.PaymentMethod = russiaPost.PaymentMethodCASHLESS
	request.Mass = order.WeightCalculate()
	resp, err := russiaPost.DefaultClient.Backlog(request)
	if err == nil {
		return err
	}

	fmt.Printf("%v \n", resp)
	return  nil
}
