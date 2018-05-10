package russiaPost

import (
	"time"
	"encoding/xml"
)

type OperationHistoryLanguage string
type MessageType string

var (
	OperationHistoryLanguageRUS OperationHistoryLanguage = "RUS"
	OperationHistoryLanguageENG OperationHistoryLanguage = "ENG"

	MessageTypeStandard MessageType = "0"
	MessageTypeOrder    MessageType = "1"
)

type (
	OperationHistory struct {
		// Идентификатор регистрируемого почтового отправления в одном из форматов:
		Barcode string `xml:"data:Barcode"`

		//Тип сообщения. Возможные значения:
		//0 - история операций для отправления;
		//1 - история операций для заказного уведомления по данному отправлению.
		MessageType MessageType `xml:"data:MessageType"`

		//Язык, на котором должны возвращаться названия операций/атрибутов и сообщения об ошибках. Допустимые значения:
		//RUS – использовать русский язык (используется по умолчанию);
		//ENG – использовать английский язык.
		Language OperationHistoryLanguage `xml:"data:Language"`
	}

	AuthorizationHeader struct {
		Login    string `xml:"data:login"`
		Password string `xml:"data:password"`
		Key      string `xml:"soapenv:mustUnderstand,attr"`
	}

	Request struct {
		OperationHistoryRequest OperationHistory    `xml:"data:OperationHistoryRequest"`
		AuthorizationHeader     AuthorizationHeader `xml:"data:AuthorizationHeader"`
	}

	Address struct {
		Index       string `xml:"Index"`
		Description string `xml:"Description"`
	}

	//Содержит финансовые данные, связанные с операцией над почтовым отправлением.
	FinanceParameters struct {
		//Сумма наложенного платежа в копейках.
		Payment int `xml:"Payment"`
		//Сумма объявленной ценности в копейках.
		Value int `xml:"Value"`
		//Сумма дополнительного тарифного сбора в копейках.
		Rate int `xml:"Rate"`
		//Сумма платы за объявленную ценность в копейках.
		InsrRate int `xml:"InsrRate"`
		//Выделенная сумма платы за пересылку воздушным транспортом из общей суммы платы за пересылку в копейках.
		AirRate int `xml:"AirRate"`
		//Общая сумма платы за пересылку наземным и воздушным транспортом в копейках.
		MassRate int `xml:"MassRate"`
		//Сумма таможенного платежа в копейках.
		CustomDuty int `xml:"CustomDuty"`
	}

	//Cодержит параметры операции над отправлением
	OperationParameters struct {
		//Содержит информацию об операции над отправлением.
		Type string    `xml:"OperType>Name"`
		//Содержит информацию об операции над отправлением.
		TypeID OperationType    `xml:"OperType>Id"`
		//Содержит информацию об атрибуте операции над отправлением.
		Attr string    `xml:"OperAttr>Name"`
		//Содержит информацию об атрибуте операции над отправлением.
		AttrID int    `xml:"OperAttr>Id"`
		//Содержит данные о дате и времени проведения операции над отправлением.
		Date time.Time `xml:"OperDate"`
	}

	UserParameters struct {
		//Содержит данные об отправителе.
		Sender    string `xml:"Sndr"`
		//Содержит данные о получателе отправления.
		Recipient string `xml:"Rcpn"`
	}

	//Содержит адресные данные места проведения операции над отправлением.
	AddressParameters struct {
		Index       string `xml:"OperationAddress>Index"`
		Description string `xml:"OperationAddress>Description"`
	}

	HistoryRecord struct {
		Finance   FinanceParameters   `xml:"FinanceParameters"`
		User      UserParameters      `xml:"UserParameters"`
		Address   AddressParameters   `xml:"AddressParameters"`
		Operation OperationParameters `xml:"OperationParameters"`
	}

	OperationHistoryResponse struct {
		XMLName              xml.Name        `xml:"getOperationHistoryResponse"`
		OperationHistoryData []HistoryRecord `xml:"OperationHistoryData>historyRecord"`
	}
)
