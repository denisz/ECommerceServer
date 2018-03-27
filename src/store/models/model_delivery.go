package models

type DeliveryMethod string
type DeliveryProvider string

const (
	// Почта России
	DeliveryProviderRussiaPost DeliveryProvider = "russiaPostal"

	// Boxberry
	DeliveryProviderBoxberry DeliveryProvider = "boxberry"

	// Курьерский способ доставки
	DeliveryMethodEMC DeliveryMethod = "emc"

	// Ускоренный способ доставки
	DeliveryMethodRapid DeliveryMethod = "rapid"

	// Обычный способ доставки
	DeliveryMethodStandart DeliveryMethod = "standart"
)

type (
	Delivery struct {
		Provider DeliveryProvider `json:"provider"`
		Method DeliveryMethod `json:"method"`
	}

)