package models

type DeliveryMethod string
type DeliveryProvider string

const (
	// Почта России
	DeliveryProviderRussiaPost DeliveryProvider = "russiapost"

	// Boxberry
	DeliveryProviderBoxberry DeliveryProvider = "boxberry"

	// CDEK
	DeliveryProviderCDEK DeliveryProvider = "cdek"

	// Байкал
	DeliveryProviderBaikal DeliveryProvider = "baikal"

	// ПЭК
	DeliveryProviderPEC DeliveryProvider = "pec"

	// Энергия
	DeliveryProviderNRG DeliveryProvider = "nrg"

	// Неизвестный способ доставки
	DeliveryMethodUnknown DeliveryMethod = ""

	// Курьерский способ доставки для почты россии
	DeliveryMethodRussiaPostEMC DeliveryMethod = "pochta_emc"

	// Ускоренный способ доставки для почты россии
	DeliveryMethodRussiaPostRapid DeliveryMethod = "pochta_rapid"

	// Обычный способ доставки для почты россии
	DeliveryMethodRussiaPostStandard DeliveryMethod = "pochta_standard"

	// Курьерский способ доставки для cdek
	DeliveryMethodCDEKEMC DeliveryMethod = "cdek_emc"

	// Ускоренный способ доставки для cdek
	DeliveryMethodCDEKRapid DeliveryMethod = "cdek_rapid"

	// Обычный способ доставки для cdek
	DeliveryMethodCDEKStandard DeliveryMethod = "cdek_standard"
)

type (
	// Доставка
	Delivery struct {
		// Повставщик услуг
		Provider DeliveryProvider `json:"provider"`
		// Способ доставки
		Method   DeliveryMethod   `json:"method"`
	}
)
