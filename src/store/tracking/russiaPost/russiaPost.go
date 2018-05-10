package russiaPost

import (
	"store/soap"
)

func TrackRussiaPost(barcode string) (*OperationHistoryResponse, error) {
	req := Request {
		OperationHistoryRequest: OperationHistory {
			Barcode: barcode,
			MessageType: MessageTypeStandard,
			Language: OperationHistoryLanguageRUS,
		},
		AuthorizationHeader: AuthorizationHeader{
			Login: "IqbsnBkvzDwjUo",
			Password: "1i8JwE1XsDRM",
			Key: "1",
		},
	}

	client, err := soap.SoapClient("http://tracking.russianpost.ru/rtm34?wsdl", false)
	if err != nil {
		return nil, err
	}

	err = client.Call("getOperationHistory", req)
	if err != nil {
		return nil, err
	}

	data := OperationHistoryResponse{}
	client.Unmarshal(&data)

	return &data, nil
}
