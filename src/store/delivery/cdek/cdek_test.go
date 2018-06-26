package cdek

import (
	"testing"
)

/**
E.g.
{
	"version":"1.0",
	"dateExecute":"2012-07-27",
	"authLogin":"098f6bcd4621d373cade4e832627b4f6",
	"secure":"396fe8e7dfd37c7c9f361bba60db0874",
	"senderCityId":"270",
	"receiverCityId":"44",
	"tariffId":"137",
	"goods":
		[
			{
				"weight":"0.3",
				"length":"10",
				"width":"7",
				"height":"5"
			},
			{
				"weight":"0.1",
				"volume":"0.1"
			}
		],
	"services": [
		{
			"id": 2,
			"param": 2000
		},
		{
			"id": 30
		}
	]
}
 */

func TestCDEK_Tariff(t *testing.T) {
	r := DestinationRequest{
		SenderCityID:   417,
		ReceiverCityID: 16197,
		TariffID:       137,
		ModeID:         ModeDeliveryDoorDoor,
		Goods: []Dimension{
			{
				Weight: 0.3,
				Length: 10,
				Width:  7,
				Height: 5,
			},
		},
		Services: []ServiceRequest{
			{
				ID:    ServiceCodeEnsure,
				Param: 2000,
			},
			{
				ID: ServiceCodeDressingRoom,
			},
		},
	}

	DebugClient.Tariff(r)
}
