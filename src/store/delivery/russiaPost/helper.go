package russiaPost

import (
	"strconv"
	"fmt"
)

func CheckValidateAddress(address *NormalizeAddress) error {
	//Код проверки должен быть: VALIDATED, OVERRIDDEN или CONFIRMED_MANUALLY.
	isValid := address.ValidationCode == ValidationCodeCONFIRMED_MANUALLY ||
		address.ValidationCode == ValidationCodeOVERRIDDEN ||
		address.ValidationCode == ValidationCodeVALIDATED

	//Код качества должен быть: GOOD, POSTAL_BOX, ON_DEMAND или UNDEF_05.
	isQuality := address.QualityCode == QualityCodeGOOD ||
		address.QualityCode == QualityCodeON_DEMAND ||
		address.QualityCode == QualityCodePOSTAL_BOX ||
		address.QualityCode == QualityCodeUNDEF_05

	if isValid && isQuality {
		return nil
	}

	return ErrAddressNotValid
}

func CheckValidatePhysical(physical *NormalizePhysical) error {
	if physical.QualityCode == QualityPhysicalCodeEDITED {
		return nil
	}
	return ErrPhysicalNotValid
}

func CheckValidatePhone(phone *NormalizePhone) error {
	if phone.QualityCode == QualityPhoneCodeGOOD ||
		phone.QualityCode == QualityPhoneCodeGOOD_REPLACED_NUMBER {
		return nil
	}
	return ErrPhoneNotValid
}

func UpdateOrderRequestWithAddress(req *OrderRequest, address *NormalizeAddress) {
	index, err := strconv.Atoi(address.Index)
	if err != nil {
		return
	}

	req.Comment = ""
	req.AddressType = address.AddressType
	req.Index = index
	req.Area = address.Area
	req.Building = address.Building
	req.Corpus = address.Corpus
	req.Hotel = address.Hotel
	req.House = address.House
	req.Letter = address.Letter
	req.Place = address.Place
	req.Region = address.Region
	req.Room = address.Room
	req.Slash = address.Slash
	req.Street = address.Street
}

func UpdateOrderRequestWithPhysical(req *OrderRequest, physical *NormalizePhysical) {
	req.GivenName = physical.Name
	req.Surname = physical.Surname
	req.MiddleName = physical.MiddleName
	req.RecipientName = fmt.Sprintf("%s %s %s", physical.Surname, physical.Name, physical.MiddleName)
}

func UpdateOrderRequestWithPhone(req *OrderRequest, phone *NormalizePhone) {
	fullNumber := fmt.Sprintf("%s%s%s",phone.PhoneCountryCode, phone.PhoneCityCode, phone.PhoneNumber)
	telAddress, err := strconv.Atoi(fullNumber)
	if err != nil {
		return
	}
	req.TelAddress = telAddress
}
