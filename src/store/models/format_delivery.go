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
		case DeliveryMethodStandard:
			buffer.WriteString(" - Стандарт")

		case DeliveryMethodEMC:
			buffer.WriteString(" - Курьерская служба")

		case DeliveryMethodRapid:
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
