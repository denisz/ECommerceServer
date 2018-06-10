package models

import (
	"bytes"
)

func (p *Delivery) Format() string {
	var buffer bytes.Buffer

	switch p.Provider {
	case DeliveryProviderRussiaPost:
		buffer.WriteString("Почта России")
		switch p.Method {
		case DeliveryMethodRussiaPostStandard:
			buffer.WriteString(" - Стандарт")

		case DeliveryMethodRussiaPostEMC:
			buffer.WriteString(" - Курьерская служба")

		case DeliveryMethodRussiaPostRapid:
			buffer.WriteString(" - Ускоренная")
		}

	case DeliveryProviderBoxberry:
		buffer.WriteString("Boxberry")

	case DeliveryProviderNRG:
		buffer.WriteString("Почтовая служба \"Энергия\"")

	case DeliveryProviderPEC:
		buffer.WriteString("Почтовая служба \"ПЭК\"")

	case DeliveryProviderBaikal:
		buffer.WriteString("Почтовая служба \"Байкал\"")
	}

	return buffer.String()
}
