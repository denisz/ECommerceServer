package api

import (
	. "store/models"
	"store/delivery/russiaPost"
	"time"
)

type ControllerForms struct {
	Controller
}

func (p *ControllerForms) FormsOrder(id int) ([]byte, error) {
	var order Order
	err := p.GetStore().
		From(NodeNamedOrders).
		One("ID", id, &order)

	if err != nil {
		return nil, err
	}

	if order.Shipment.Provider == DeliveryProviderRussiaPost {
		switch order.Shipment.Method {
		case DeliveryMethodEMC:
			return russiaPost.DefaultClient.FormsE1(order.Shipment.ExternalNumber, time.Now())
		case DeliveryMethodRapid:
			return russiaPost.DefaultClient.FormsE1(order.Shipment.ExternalNumber, time.Now())
		case DeliveryMethodStandard:
			return russiaPost.DefaultClient.FormsF7(order.Shipment.ExternalNumber, time.Now())
		}

	}

	return nil, ErrNotSupportedProvider
}


