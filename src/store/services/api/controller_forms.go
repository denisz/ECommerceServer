package api

import (
	"store/delivery/russiaPost"
	"time"
	. "store/models"
	"fmt"
	"strings"
)

type ControllerForms struct {
	Controller
}

//формы заказа
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
			return russiaPost.DefaultClient.FormsEMS(order.Shipment.ExternalNumber, time.Now())
		case DeliveryMethodRapid:
			return russiaPost.DefaultClient.FormsF7(order.Shipment.ExternalNumber, time.Now())
		case DeliveryMethodStandard:
			return russiaPost.DefaultClient.FormsF7(order.Shipment.ExternalNumber, time.Now())
		}

	}

	return nil, ErrNotSupportedProvider
}


// Формы партии
func(p *ControllerForms) FormsBatch(id int) ([]byte, error) {
	var batch Batch
	err := p.GetStore().
		From(NodeNamedBatches).
		One("ID", id, &batch)

	if err != nil {
		return nil, err
	}

	if batch.Provider == DeliveryProviderRussiaPost {
		if len(batch.PayloadRussiaPost) > 0 {
			if len(batch.PayloadRussiaPost) > 1 {
				fmt.Printf("russia post has several batch %s", strings.Join(batch.PayloadRussiaPost, ","))
			}

			batchName := batch.PayloadRussiaPost[0]
			zip, err := russiaPost.DefaultClient.FormsPacket(batchName)
			if err != nil {
				return nil, err
			}
			return russiaPost.ReadZipForms(zip)
		}
	}

	return nil, ErrNotSupportedProvider
}

