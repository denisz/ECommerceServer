package models

type DeliveryMethod string
type DeliveryProvider string

const (
	// Почта России
	DeliveryProviderRussiaPost DeliveryProvider = "russiapost"

	// Boxberry
	DeliveryProviderBoxberry DeliveryProvider = "boxberry"

	// Байкал
	DeliveryProviderBaikal DeliveryProvider = "baikal"

	// ПЭК
	DeliveryProviderPEC DeliveryProvider = "pec"

	// Энергия
	DeliveryProviderNRG DeliveryProvider = "nrg"

	// Курьерский способ доставки
	DeliveryMethodEMC DeliveryMethod = "emc"

	// Ускоренный способ доставки
	DeliveryMethodRapid DeliveryMethod = "rapid"

	// Обычный способ доставки
	DeliveryMethodStandard DeliveryMethod = "standard"
)

type (
	//Доставка
	Delivery struct {
		Provider DeliveryProvider `json:"provider"`
		Method   DeliveryMethod   `json:"method"`
	}

	//Стоимость доставки
	DeliveryPayment struct {
		Price int `json:"price"`
	}
)
