package russiaPost

import "strconv"

func CheckValidateAddress(nAddress *NormalizeAddress) error {
	//Код проверки должен быть: VALIDATED, OVERRIDDEN или CONFIRMED_MANUALLY.
	isValid := nAddress.ValidationCode == ValidationCodeCONFIRMED_MANUALLY ||
		nAddress.ValidationCode == ValidationCodeOVERRIDDEN ||
		nAddress.ValidationCode == ValidationCodeVALIDATED

	//Код качества должен быть: GOOD, POSTAL_BOX, ON_DEMAND или UNDEF_05.
	isQuality := nAddress.QualityCode == QualityCodeGOOD ||
		nAddress.QualityCode == QualityCodeON_DEMAND ||
		nAddress.QualityCode == QualityCodePOSTAL_BOX ||
		nAddress.QualityCode == QualityCodeUNDEF_05

	if isValid && isQuality {
		return nil
	}

	return ErrAddressNotValid
}

func CreateOrderRequestWithAddress(address *NormalizeAddress) OrderRequest {
	index, err := strconv.Atoi(address.Index)
	if err != nil {
		return OrderRequest{}
	}

	return OrderRequest{
		Comment:     "",
		AddressType: address.AddressType,
		Area:        address.Area,
		Building:    address.Building,
		Corpus:      address.Corpus,
		Hotel:       address.Hotel,
		House:       address.House,
		Letter:      address.Letter,
		Index:       index,
		Place:       address.Place,
		Region:      address.Region,
		Room:        address.Room,
		Slash:       address.Slash,
		Street:      address.Street,
	}
}
