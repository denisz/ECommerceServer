package pochta


type AddressType string
type EnvelopeType string
type MailCategory string
type PaymentMethod string
type MailType string

const (
	//Стандартный (улица, дом, квартира)
	AddressTypeDefault AddressType = "DEFAULT"

	//Абонентский ящик
	AddressTypePO_BOX AddressType = "PO_BOX"

	//До востребования
	AddressTypeDEMAND AddressType = "DEMAND"

	//229мм x 324мм
	EnvelopeTypeC4 EnvelopeType = "C4"

	//162мм x 229мм
	EnvelopeTypeC5 EnvelopeType = "C5"

	//220мм x 110мм
	EnvelopeTypeDL EnvelopeType = "DL"

	//148мм x 105мм
	EnvelopeTypeA6 EnvelopeType = "A6"

	//Простое
	MailCategorySIMPLE MailCategory = "SIMPLE"

	//Заказное
	MailCategoryORDERED MailCategory = "ORDERED"

	//Обыкновенное
	MailCategoryORDINARY MailCategory = "ORDINARY"

	//С объявленной ценностью
	MailCategoryWITH_DECLARED_VALUE MailCategory = "WITH_DECLARED_VALUE"

	//С объявленной ценностью и наложенным платежом
	MailCategoryWITH_DECLARED_VALUE_AND_CASH_ON_DELIVERY MailCategory = "WITH_DECLARED_VALUE_AND_CASH_ON_DELIVERY"

	//Комбинированное
	MailCategoryCOMBINED MailCategory = "COMBINED"

	//Безналичный расчет
	PaymentMethodCASHLESS PaymentMethod = "CASHLESS"

	//Оплата марками
	PaymentMethodSTAMP	PaymentMethod = "STAMP"

	//Франкирование
	PaymentMethodFRANKING PaymentMethod = "FRANKING"

	//Посылка "нестандартная"
	MailTypePOSTAL_PARCEL MailType = "POSTAL_PARCEL"

	//Посылка "онлайн"
	MailTypeONLINE_PARCEL MailType = "ONLINE_PARCEL"

	//Курьер "онлайн"
	MailTypeONLINE_COURIER MailType = "ONLINE_COURIER"

	//Отправление  EMS
	MailTypeEMS MailType = "EMS"

	//EMS оптимальное
	MailTypeEMS_OPTIMAL MailType = "EMS_OPTIMAL"

	//EMS РТ
	MailTypeEMS_RT MailType = "EMS_RT"

	//Письмо
	MailTypeLETTER MailType = "LETTER"

	//Бандероль
	MailTypeBANDEROL MailType = "BANDEROL"

	//Бизнес курьер
	MailTypeBUSINESS_COURIER MailType = "BUSINESS_COURIER"

	//Бизнес курьер экпресс
	MailTypeBUSINESS_COURIER_ES MailType = "BUSINESS_COURIER_ES"

	//Посылка 1-го класса
	MailTypePARCEL_CLASS_1 MailType = "PARCEL_CLASS_1"

	//Комбинированное
	MailTypeCOMBINED MailType = "COMBINED"
)

type (

	Dimension struct {
		//Линейная ширина (сантиметры)
		Width int `json:"width"`
		//Линейная высота (сантиметры)
		Height int `json:"height"`
		//Линейная длина (сантиметры)
		Length int `json:"length"`
	}

	OrderRequest struct {
		//Тип адреса
		AddressTypeTo AddressType `json:"address-type-to"`
		//Район
		AreaTo string `json:"area-to"`
		//Отправитель на посылке/название брэнда
		BrandName string `json:"brand-name"`
		//Часть здания: Строение
		BuildingTo string `json:"building-to"`
		//Комментарий:Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		Comment string `json:"comment"`
		//Часть здания: Корпус
		CorpusTo string `json:"corpus-to"`
		//Отметка 'Курьер'
		Courier bool `json:"courier"`
		//Линейные размеры
		Dimension Dimension `json:"dimension"`
		//Тип конверта - ГОСТ Р 51506-99.
		EnvelopeType EnvelopeType `json:"envelope-type"`
		//Установлена ли отметка 'Осторожно/Хрупкое'?
		Fragile bool `json:"fragile"`
		//Имя получателя
		GivenName string `json:"given-name"`
		//Отчество получателя
		MiddleName string `json:"middle-name"`
		//Фамилия получателя
		Surname string `json:"surname"`
		//Название гостиницы
		HotelTo string `json:"hotel-to"`
		//Часть адреса: Номер здания
		HouseTo string `json:"house-to"`
		//Почтовый индекс
		IndexTo int `json:"index-to"`
		//Сумма объявленной ценности (копейки)
		InsrValue int `json:"insr-value"`
		//Часть здания: Литера
		LetterTo string `json:"letter-to"`
		//Микрорайон
		LocationTo string `json:"location-to"`
		//Категория РПО
		MailCategory MailCategory `json:"mail-category"`
		//Код страны Россия: 643
		MailDirect int `json:"mail-direct"`
		//Вид РПО
		MailType MailType `json:"mail-type"`
		//Отметка 'Ручной ввод адреса'
		ManualAddressInput bool `json:"manual-address-input"`
		//Вес РПО (в граммах)
		Mass int `json:"mass"`
		//Номер для а/я, войсковая часть, войсковая часть ЮЯ, полевая почта
		NumAddressTypeTo string `json:"num-address-type-to"`
		//Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		OrderNum string `json:"order-num"`
		//Сумма наложенного платежа (копейки)
		Payment int `json:"payment"`
		//Способ оплаты.
		PaymentMethod PaymentMethod `json:"payment-method"`
		//Населенный пункт
		PlaceTo string `json:"place-to"`
		//Индекс места приема
		PostOfficeCode string `json:"postoffice-code"`
		//Наименование получателя одной строкой (ФИО, наименование организации)
		RecipientName string `json:"recipient-name"`
		//Область, регион
		RegionTo string `json:"region-to"`
		//Часть здания: Номер помещения
		RoomTo string `json:"room-to"`
		//Часть здания: Дробь
		SlashTo string `json:"slash-to"`
		//Признак услуги SMS уведомления
		SMSNoticeRecipient int `json:"sms-notice-recipient"`
		//Часть адреса: Улица
		StreetTo string `json:"street-to"`
		//Телефон получателя (может быть обязательным для некоторых типов отправлений)
		TelAddress int `json:"tel-address"`
		//Отметка 'С заказным уведомлением'
		WithOrderOfNotice bool `json:"with-order-of-notice"`
		//Отметка 'С простым уведомлением'
		WithSimpleNotice bool `json:"with-simple-notice"`
		//Отметка 'Без разряда'
		WoMailRank bool `json:"wo-mail-rank"`
	}

	OrderError struct {
		//Список кодов ошибок
		Codes []struct {
			Code string `json:"code"`
			Details string `json:"details"`
		} `json:"error-codes"`
		//Индекс в исходном массиве
		Position int `json:"position"`
	}

	OrderResponse struct {
		//Список ошибок
		Errors []OrderError `json:"errors"`
		//Список идентификаторов успешно обработанных сущностей
		Ids []int `json:"result-ids"`
	}

	DestinationRequest struct {
		//Отметка 'Курьер'
		Courier bool `json:"courier"`
		//Объявленная ценность
		DeclareValue int `json:"declared-value"`
		//Линейные размеры
		Dimension Dimension `json:"dimension"`
		//Установлена ли отметка 'Осторожно/Хрупкое'?
		Fragile bool `json:"fragile"`
		//Почтовый индекс объекта почтовой связи места приема
		IndexFrom string `json:"index-from"`
		//Почтовый индекс объекта почтовой связи места назначения
		IndexTo string `json:"index-to"`
		//Категория РПО
		MailCategory MailCategory `json:"mail-category"`
		//Вид РПО
		MailType MailType `json:"mail-type"`
		//Вес РПО (в граммах)
		Mass int `json:"mass"`
		//Способ оплаты.
		PaymentMethod PaymentMethod `json:"payment-method"`
		//Отметка 'С заказным уведомлением'
		WithOrderOfNotice bool `json:"with-order-of-notice"`
		//Отметка 'С простым уведомлением'
		WithSimpleNotice bool `json:"with-simple-notice"`
	}

	DestinationRate struct {
		//Тариф без НДС (коп)
		Rate int `json:"rate"`
		//НДС (коп)
		Vat int `json:"vat"`
	}

	DeliveryTime struct {
		//Максимальное время доставки (дни)
		MaxDays int `json:"max-days"`
		//Минимальное время доставки (дни)
		MinDays int `json:"min-days"`
	}

	DestinationResponse struct {
		//Плата за Авиа-пересылку (коп)
		AviaRate DestinationRate `json:"avia-rate"`
		//Надбавка за отметку 'Осторожно/Хрупкое'
		FragileRate DestinationRate `json:"fragile-rate"`
		//Плата за пересылку (коп)
		GroundRate DestinationRate `json:"ground-rate"`
		//Плата за объявленную ценность (коп)
		InsuranceRate DestinationRate `json:"insurance-rate"`
		//Надбавка за уведомление о вручении
		NoticeRate DestinationRate `json:"notice-rate"`
		//Надбавка за негабарит при весе более 10кг
		OversizeRate DestinationRate `json:"oversize-rate"`
		//Время доставки
		DeliveryTime DeliveryTime `json:"delivery-time"`
		//Плата всего (коп)
		TotalRate int `json:"total-rate"`
		//Всего НДС (коп)
		TotalVat int `json:"total-vat"`
	}
)