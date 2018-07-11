package api

import (
	. "store/models"
	"store/delivery/russiaPost"
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

func NormalizePhysicalForRussiaPost(address *Address) (*russiaPost.NormalizePhysical, error) {
	req := russiaPost.NormalizePhysicalRequest{
		ID:             "Physical 1",
		OriginalString: address.Name,
	}

	return russiaPost.DefaultClient.NormalizePhysical(req)
}

func NormalizePhoneForRussiaPost(address *Address) (*russiaPost.NormalizePhone, error) {
	req := russiaPost.NormalizePhoneRequest{
		ID:             "Phone 1",
		OriginalString: address.Phone,
	}

	return russiaPost.DefaultClient.NormalizePhone(req)
}

func CheckValidAddress(address *Address) error {
	nAddress, err := NormalizeAddressForRussiaPost(address)
	if err != nil {
		return err
	}

	if nAddress.Index != address.PostalCode {
		return ErrILLEGAL_INDEX_TO
	}

	err = russiaPost.CheckValidateAddress(nAddress)
	if err != nil {
		return err
	}

	nPhysical, err := NormalizePhysicalForRussiaPost(address)
	if err != nil {
		return err
	}

	err = russiaPost.CheckValidatePhysical(nPhysical)
	if err != nil {
		return err
	}

	nPhone, err := NormalizePhoneForRussiaPost(address)
	if err != nil {
		return err
	}

	err = russiaPost.CheckValidatePhone(nPhone)
	if err != nil {
		return err
	}

	return nil
}

func CreateOrderInToRussiaPost(order *Order) (*russiaPost.Order, error) {
	if order.Delivery.Provider != DeliveryProviderRussiaPost {
		return nil, ErrNotSupportedProvider
	}

	address, err := NormalizeAddressForRussiaPost(order.Address)
	if err != nil {
		return nil, err
	}

	phone, err := NormalizePhoneForRussiaPost(order.Address)
	if err != nil {
		return nil, err
	}

	physical, err := NormalizePhysicalForRussiaPost(order.Address)
	if err != nil {
		return nil, err
	}

	request := russiaPost.OrderRequest{
		PostOfficeCode: "430005",
		OrderNum:     order.Invoice,
		BrandName:    "DarkWaters",
		MailDirect:   643,
		MailCategory: russiaPost.MailCategoryORDINARY,
		Mass:         order.WeightCalculate().Gram(),
	}

	switch order.Delivery.Method {
	case DeliveryMethodRussiaPostEMC:
		request.MailType = russiaPost.MailTypeBUSINESS_COURIER
	case DeliveryMethodRussiaPostRapid:
		request.MailType = russiaPost.MailTypePARCEL_CLASS_1
	case DeliveryMethodRussiaPostStandard:
		request.MailType = russiaPost.MailTypePOSTAL_PARCEL
	}

	russiaPost.UpdateOrderRequestWithPhone(&request, phone)
	russiaPost.UpdateOrderRequestWithAddress(&request, address)
	russiaPost.UpdateOrderRequestWithPhysical(&request, physical)

	orderID, err := russiaPost.DefaultClient.CreateBacklog(request)
	if err != nil {
		return nil,err
	}

	return russiaPost.DefaultClient.GetBacklog(orderID)
}

//func CreateOrderInCDEKProvider(order *Order) (*cdek.OrderCreateRequest, error) {
//
//}

func DefaultMethodDeliveryForProvider(provider DeliveryProvider) DeliveryMethod {
	switch provider {
	case DeliveryProviderRussiaPost:
		return DeliveryMethodRussiaPostRapid
		//return DeliveryMethodRussiaPostStandard
	case DeliveryProviderCDEK:
		return DeliveryMethodCDEKRapid
		//return DeliveryMethodCDEKStandard
	default:
		return DeliveryMethodUnknown
	}
}