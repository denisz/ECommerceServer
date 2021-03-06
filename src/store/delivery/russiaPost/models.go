package russiaPost

type AddressType string
type EnvelopeType string
type MailCategory string
type PaymentMethod string
type MailType string
type QualityCode string
type QualityPhysicalCode string
type QualityPhoneCode string
type ValidationCode string
type BatchStatusCode string

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
	PaymentMethodSTAMP PaymentMethod = "STAMP"

	//Франкирование
	PaymentMethodFRANKING PaymentMethod = "FRANKING"

	//Посылка "нестандартная"
	MailTypePOSTAL_PARCEL MailType = "POSTAL_PARCEL"

	//Посылка "онлайн"
	MailTypeONLINE_PARCEL MailType = "ONLINE_PARCEL"

	//Курьер "онлайн"
	MailTypeONLINE_COURIER MailType = "ONLINE_COURIER"

	//Отправление EMS (не работает)
	MailTypeEMS MailType = "EMS"

	//EMS оптимальное (не работает)
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

	//Пригоден для почтовой рассылки
	QualityCodeGOOD QualityCode = "GOOD"

	//До востребования
	QualityCodeON_DEMAND QualityCode = "ON_DEMAND"

	//Абонентский ящик
	QualityCodePOSTAL_BOX QualityCode = "POSTAL_BOX"

	//Не определен регион
	QualityCodeUNDEF_01 QualityCode = "UNDEF_01"

	//Не определен город или населенный пункт
	QualityCodeUNDEF_02 QualityCode = "UNDEF_02"

	//Не определена улица
	QualityCodeUNDEF_03 QualityCode = "UNDEF_03"

	//Не определен номер дома
	QualityCodeUNDEF_04 QualityCode = "UNDEF_04"

	//Не определена квартира/офис
	QualityCodeUNDEF_05 QualityCode = "UNDEF_05"

	//Не определен
	QualityCodeUNDEF_06 QualityCode = "UNDEF_06"

	//Иностранный адрес
	QualityCodeUNDEF_07 QualityCode = "UNDEF_07"

	//Подтверждено контролером
	ValidationCodeCONFIRMED_MANUALLY ValidationCode = "CONFIRMED_MANUALLY"

	//Уверенное распознавание
	ValidationCodeVALIDATED ValidationCode = "VALIDATED"

	//Распознан: адрес был перезаписан в справочнике
	ValidationCodeOVERRIDDEN ValidationCode = "OVERRIDDEN"

	//На проверку, неразобранные части
	ValidationCodeNOT_VALIDATED_HAS_UNPARSED_PARTS ValidationCode = "NOT_VALIDATED_HAS_UNPARSED_PARTS"

	//На проверку, предположение
	ValidationCodeNOT_VALIDATED_HAS_ASSUMPTION ValidationCode = "NOT_VALIDATED_HAS_ASSUMPTION"

	//На проверку, нет основных частей
	ValidationCodeNOT_VALIDATED_HAS_NO_MAIN_POINTS ValidationCode = "NOT_VALIDATED_HAS_NO_MAIN_POINTS"

	//На проверку, предположение по улице
	ValidationCodeNOT_VALIDATED_HAS_NUMBER_STREET_ASSUMPTION ValidationCode = "NOT_VALIDATED_HAS_NUMBER_STREET_ASSUMPTION"

	//На проверку, нет в КЛАДР
	ValidationCodeNOT_VALIDATED_HAS_NO_KLADR_RECORD ValidationCode = "NOT_VALIDATED_HAS_NO_KLADR_RECORD"

	//На проверку, нет улицы или населенного пункта
	ValidationCodeNOT_VALIDATED_HOUSE_WITHOUT_STREET_OR_NP ValidationCode = "NOT_VALIDATED_HOUSE_WITHOUT_STREET_OR_NP"

	//На проверку, нет дома
	ValidationCodeNOT_VALIDATED_HOUSE_EXTENSION_WITHOUT_HOUSE ValidationCode = "NOT_VALIDATED_HOUSE_EXTENSION_WITHOUT_HOUSE"

	//На проверку, неоднозначность
	ValidationCodeNOT_VALIDATED_HAS_AMBI ValidationCode = "NOT_VALIDATED_HAS_AMBI"

	//На проверку, большой номер дома
	ValidationCodeNOT_VALIDATED_EXCEDED_HOUSE_NUMBER ValidationCode = "NOT_VALIDATED_EXCEDED_HOUSE_NUMBER"

	//На проверку, некорректный дом
	ValidationCodeNOT_VALIDATED_INCORRECT_HOUSE ValidationCode = "NOT_VALIDATED_INCORRECT_HOUSE"

	//На проверку, некорректное расширение дома
	ValidationCodeNOT_VALIDATED_INCORRECT_HOUSE_EXTENSION ValidationCode = "NOT_VALIDATED_INCORRECT_HOUSE_EXTENSION"

	//Иностранный адрес
	ValidationCodeNOT_VALIDATED_FOREIGN ValidationCode = "NOT_VALIDATED_FOREIGN"

	//На проверку, не по справочнику
	ValidationCodeNOT_VALIDATED_DICTIONARY ValidationCode = "NOT_VALIDATED_DICTIONARY"

	//Подтверждено контролером
	QualityPhysicalCodeCONFIRMED_MANUALLY QualityPhysicalCode = "CONFIRMED_MANUALLY"

	//Правильное значение
	QualityPhysicalCodeEDITED QualityPhysicalCode = "EDITED"

	//Сомнительное значение
	QualityPhysicalCodeNOT_SURE QualityPhysicalCode = "NOT_SURE"

	//Подтверждено контролером
	QualityPhoneCodeCONFIRMED_MANUALLY QualityPhoneCode = "CONFIRMED_MANUALLY"

	//Корректный телефонный номер
	QualityPhoneCodeGOOD QualityPhoneCode = "GOOD"

	//Изменен код телефонного номера
	QualityPhoneCodeGOOD_REPLACED_CODE QualityPhoneCode = "GOOD_REPLACED_CODE"

	//Изменен номер телефона
	QualityPhoneCodeGOOD_REPLACED_NUMBER QualityPhoneCode = "GOOD_REPLACED_NUMBER"

	//Изменен код и номер телефона
	QualityPhoneCodeGOOD_REPLACED_CODE_NUMBER QualityPhoneCode = "GOOD_REPLACED_CODE_NUMBER"

	//Конфликт по городу
	QualityPhoneCodeGOOD_CITY_CONFLICT QualityPhoneCode = "GOOD_CITY_CONFLICT"

	//Конфликт по региону
	QualityPhoneCodeGOOD_REGION_CONFLICT QualityPhoneCode = "GOOD_REGION_CONFLICT"

	//Иностранный телефонный номер
	QualityPhoneCodeFOREIGN QualityPhoneCode = "FOREIGN"

	//Неоднозначный код телефонного номера
	QualityPhoneCodeCODE_AMBI QualityPhoneCode = "CODE_AMBI"

	//Пустой телефонный номер
	QualityPhoneCodeEMPTY QualityPhoneCode = "EMPTY"

	//Телефонный номер содержит некорректные символы
	QualityPhoneCodeGARBAGE QualityPhoneCode = "GARBAGE"

	//Восстановлен город в телефонном номере
	QualityPhoneCodeGOOD_CITY QualityPhoneCode = "GOOD_CITY"

	//Запись содержит более одного телефона
	QualityPhoneCodeGOOD_EXTRA_PHONE QualityPhoneCode = "GOOD_EXTRA_PHONE"

	//Телефон не может быть распознан
	QualityPhoneCodeUNDEF QualityPhoneCode = "UNDEF"

	//Телефон не может быть распознан
	QualityPhoneCodeINCORRECT_DATA QualityPhoneCode = "INCORRECT_DATA"

	//Партия создана
	BatchStatusCodeCREATED BatchStatusCode = "CREATED"

	//Партия в процессе приема, редактирование запрещено
	BatchStatusCodeFROZEN BatchStatusCode = "FROZEN"

	//Партия принята в отделении связи
	BatchStatusCodeACCEPTED BatchStatusCode = "ACCEPTED"

	//По заказам в партии существуют данные в сервисе трекинга
	BatchStatusCodeSENT BatchStatusCode = "SENT"

	//Партия находится в архиве
	BatchStatusCodeARCHIVED BatchStatusCode = "ARCHIVED"
)

type (
	//Размеры
	Dimension struct {
		//Линейная ширина (сантиметры)
		Width int `json:"width"`
		//Линейная высота (сантиметры)
		Height int `json:"height"`
		//Линейная длина (сантиметры)
		Length int `json:"length"`
	}

	//Время доставки
	DeliveryTime struct {
		//Максимальное время доставки (дни)
		MaxDays int `json:"max-days"`
		//Минимальное время доставки (дни)
		MinDays int `json:"min-days"`
	}

	//Запрос создания заказа
	OrderRequest struct {
		//Тип адреса
		AddressType AddressType `json:"address-type-to,omitempty"`
		//Район
		Area string `json:"area-to,omitempty"`
		//Отправитель на посылке/название брэнда
		BrandName string `json:"-"`
		//Часть здания: Строение
		Building string `json:"building-to,omitempty"`
		//Комментарий:Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		Comment string `json:"comment,omitempty"`
		//Часть здания: Корпус
		Corpus string `json:"corpus-to,omitempty"`
		//Отметка 'Курьер'
		Courier bool `json:"courier,omitempty"`
		//Линейные размеры
		Dimension Dimension `json:"dimension,omitempty"`
		//Тип конверта - ГОСТ Р 51506-99.
		EnvelopeType EnvelopeType `json:"envelope-type,omitempty"`
		//Установлена ли отметка 'Осторожно/Хрупкое'?
		Fragile bool `json:"fragile,omitempty"`
		//Имя получателя
		GivenName string `json:"given-name,omitempty"`
		//Отчество получателя
		MiddleName string `json:"middle-name,omitempty"`
		//Фамилия получателя
		Surname string `json:"surname,omitempty"`
		//Название гостиницы
		Hotel string `json:"hotel-to,omitempty"`
		//Часть адреса: Номер здания
		House string `json:"house-to,omitempty"`
		//Почтовый индекс
		Index int `json:"index-to,omitempty"`
		//Сумма объявленной ценности (копейки)
		InsrValue int `json:"insr-value,omitempty"`
		//Часть здания: Литера
		Letter string `json:"letter-to,omitempty"`
		//Микрорайон
		Location string `json:"location-to,omitempty"`
		//Категория РПО
		MailCategory MailCategory `json:"mail-category,omitempty"`
		//Код страны Россия: 643
		MailDirect int `json:"mail-direct,omitempty"`
		//Вид РПО
		MailType MailType `json:"mail-type,omitempty"`
		//Отметка 'Ручной ввод адреса'
		ManualAddressInput bool `json:"manual-address-input,omitempty"`
		//Вес РПО (в граммах)
		Mass int `json:"mass,omitempty"`
		//Номер для а/я, войсковая часть, войсковая часть ЮЯ, полевая почта
		NumAddressTypeTo string `json:"num-address-type-to,omitempty"`
		//Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		OrderNum string `json:"order-num,omitempty"`
		//Сумма наложенного платежа (копейки)
		Payment int `json:"payment,omitempty"`
		//Способ оплаты.
		PaymentMethod PaymentMethod `json:"payment-method,omitempty"`
		//Населенный пункт
		Place string `json:"place-to,omitempty"`
		//Индекс места приема
		PostOfficeCode string `json:"postoffice-code,omitempty"`
		//Наименование получателя одной строкой (ФИО, наименование организации)
		RecipientName string `json:"recipient-name,omitempty"`
		//Область, регион
		Region string `json:"region-to,omitempty"`
		//Часть здания: Номер помещения
		Room string `json:"room-to,omitempty"`
		//Часть здания: Дробь
		Slash string `json:"slash-to,omitempty"`
		//Признак услуги SMS уведомления
		SMSNoticeRecipient int `json:"sms-notice-recipient,omitempty"`
		//Часть адреса: Улица
		Street string `json:"street-to,omitempty"`
		//Телефон получателя (может быть обязательным для некоторых типов отправлений)
		TelAddress int `json:"tel-address,omitempty"`
		//Отметка 'С заказным уведомлением'
		WithOrderOfNotice bool `json:"with-order-of-notice,omitempty"`
		//Отметка 'С простым уведомлением'
		WithSimpleNotice bool `json:"with-simple-notice,omitempty"`
		//Отметка 'Без разряда'
		WoMailRank bool `json:"wo-mail-rank,omitempty"`
	}

	//Заказ
	Order struct {
		ID int `json:"id"`
		//Тип адреса
		AddressType AddressType `json:"address-type-to,omitempty"`
		//Район
		Area string `json:"area-to,omitempty"`
		//Отправитель на посылке/название брэнда
		BrandName string `json:"-"`
		//Часть здания: Строение
		Building string `json:"building-to,omitempty"`
		//Комментарий:Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		Comment string `json:"comment,omitempty"`
		//Часть здания: Корпус
		Corpus string `json:"corpus-to,omitempty"`
		//Отметка 'Курьер'
		Courier bool `json:"courier,omitempty"`
		//Линейные размеры
		Dimension Dimension `json:"dimension,omitempty"`
		//Тип конверта - ГОСТ Р 51506-99.
		EnvelopeType EnvelopeType `json:"envelope-type,omitempty"`
		//Установлена ли отметка 'Осторожно/Хрупкое'?
		Fragile bool `json:"fragile,omitempty"`
		//Имя получателя
		GivenName string `json:"given-name,omitempty"`
		//Отчество получателя
		MiddleName string `json:"middle-name,omitempty"`
		//Фамилия получателя
		Surname string `json:"surname,omitempty"`
		//Название гостиницы
		Hotel string `json:"hotel-to,omitempty"`
		//Часть адреса: Номер здания
		House string `json:"house-to,omitempty"`
		//Почтовый индекс
		Index int `json:"index-to,omitempty"`
		//Сумма объявленной ценности (копейки)
		InsrValue int `json:"insr-value,omitempty"`
		//Часть здания: Литера
		Letter string `json:"letter-to,omitempty"`
		//Микрорайон
		Location string `json:"location-to,omitempty"`
		//Категория РПО
		MailCategory MailCategory `json:"mail-category,omitempty"`
		//Код страны Россия: 643
		MailDirect int `json:"mail-direct,omitempty"`
		//Вид РПО
		MailType MailType `json:"mail-type,omitempty"`
		//Отметка 'Ручной ввод адреса'
		ManualAddressInput bool `json:"manual-address-input,omitempty"`
		//Вес РПО (в граммах)
		Mass int `json:"mass,omitempty"`
		//Номер для а/я, войсковая часть, войсковая часть ЮЯ, полевая почта
		NumAddressTypeTo string `json:"num-address-type-to,omitempty"`
		//Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		OrderNum string `json:"order-num,omitempty"`
		//Сумма наложенного платежа (копейки)
		Payment int `json:"payment,omitempty"`
		//Способ оплаты.
		PaymentMethod PaymentMethod `json:"payment-method,omitempty"`
		//Населенный пункт
		Place string `json:"place-to,omitempty"`
		//Индекс места приема
		PostOfficeCode string `json:"postoffice-code,omitempty"`
		//Наименование получателя одной строкой (ФИО, наименование организации)
		RecipientName string `json:"recipient-name,omitempty"`
		//Область, регион
		Region string `json:"region-to,omitempty"`
		//Часть здания: Номер помещения
		Room string `json:"room-to,omitempty"`
		//Часть здания: Дробь
		Slash string `json:"slash-to,omitempty"`
		//Признак услуги SMS уведомления
		SMSNoticeRecipient int `json:"sms-notice-recipient,omitempty"`
		//Часть адреса: Улица
		Street string `json:"street-to,omitempty"`
		//Телефон получателя (может быть обязательным для некоторых типов отправлений)
		TelAddress int `json:"tel-address,omitempty"`
		//Отметка 'С заказным уведомлением'
		WithOrderOfNotice bool `json:"with-order-of-notice,omitempty"`
		//Отметка 'С простым уведомлением'
		WithSimpleNotice bool `json:"with-simple-notice,omitempty"`
		//Отметка 'Без разряда'
		WoMailRank bool `json:"wo-mail-rank,omitempty"`
		//Время доставки
		DeliveryTime DeliveryTime `json:"delivery-time"`
		//ШПИ
		Barcode string `json:"barcode"`
		//Заказ удален?
		IsDelete bool `json:"is-deleted"`
		//Плата всего (коп)
		TotalRate int `json:"total-rate-wo-vat"`
		//Всего НДС (коп)
		TotalVat int `json:"total-vat"`
		//Версия заказа
		Version int `json:"version"`
	}

	//Партия
	Batch struct {
		//Номер партии
		BatchName string `json:"batch-name,omitempty"`
		//Статус партии
		BatchStatus BatchStatusCode `json:"batch-status,omitempty"`
		//Дата обновления статуса
		BatchStatusDate string `json:"batch-status-date,omitempty"`
		//Способ оплаты уведомления о вручении РПО. См. Способ оплаты
		DeliveryNoticePaymentMethod PaymentMethod `json:"delivery-notice-payment-method,omitempty"`
		//Номер документа для сдачи партии
		ListNumber int `json:"list-number,omitempty"`
		//Дата документа для сдачи партии
		ListNumberDate string `json:"list-number-date,omitempty"`
		//Категория РПО. См. Категория РПО
		MailCategory MailCategory `json:"mail-category,omitempty"`
		//Категория РПО (текст)
		MailCategoryText string `json:"mail-category-text,omitempty"`
		//Код разряда почтового отправления
		MailRank string `json:"mail-rank,omitempty"`
		//Вид РПО. См. Вид РПО
		MailType MailType `json:"mail-type,omitempty"`
		//Вид РПО (текст)
		MailTypeText string `json:"mail-type-text,omitempty"`
		//Способ оплаты. См. Способ оплаты
		PaymentMethod PaymentMethod `json:"payment-method,omitempty"`
		//Отметки внутренних и международных отправлений
		Postmarks []string `json:"postmarks,omitempty"`
		//Адрес места приема
		PostOfficeAddress string `json:"postoffice-address,omitempty"`
		//Индекс места приема
		PostOfficeCode string `json:"postoffice-code,omitempty"`
		//Наименование места приема
		PostOfficeName string `json:"postoffice-name,omitempty"`
		//Сумма платы за авиа пересылку в копейках, без НДС
		ShipmentAviaRate int `json:"shipment-avia-rate-sum,omitempty"`
		//НДС от суммы платы за авиа пересылку в копейках
		ShipmentAviaRateVat int `json:"shipment-avia-rate-vat-sum,omitempty"`
		//Количество заказов в партии
		ShipmentCount int `json:"shipment-count,omitempty"`
		//Сумма платы за наземную пересылку в копейках, без НДС
		ShipmentGroundRate int `json:"shipment-ground-rate-sum,omitempty"`
		//НДС от суммы платы за наземную пересылку в копейках
		ShipmentGroundRateVat int `json:"shipment-ground-rate-vat-sum,omitempty"`
		//Сумма платы за объявленную ценность в копейках, без НДС
		ShipmentInsureRate int `json:"shipment-insure-rate-sum,omitempty"`
		//НДС от суммы платы за объявленную ценность в копейках
		ShipmentInsureRateVat int `json:"shipment-insure-rate-vat-sum,omitempty"`
		//Общий вес в граммах
		ShipmentMass int `json:"shipment-mass,omitempty"`
		//Сумма платы за пересылку в копейках, без НДС
		ShipmentMassRate int `json:"shipment-mass-rate-sum,omitempty"`
		//НДС от суммы платы за пересылку в копейках
		ShipmentMassRateVat int `json:"shipment-mass-rate-vat-sum,omitempty"`
		//Сумма надбавки за уведомление о вручении в копейках
		ShipmentNoticeRate int `json:"shipment-notice-rate-sum,omitempty"`
		//НДС от суммы платы за смс нотификацию в копейках
		ShipmentNoticeRateVat int `json:"shipment-notice-rate-vat-sum,omitempty"`
		//Сумма платы за смс нотификацию в копейках, без НДС
		ShipmentSmsNoticeRate int `json:"shipment-sms-notice-rate-sum,omitempty"`
		//НДС от суммы платы за смс нотификацию в копейках
		ShipmentSmsNoticeRateVar int `json:"shipment-sms-notice-rate-vat-sum,omitempty"`
		//Категория уведомления о вручении РПО. См. Категория уведомления о вручении РПО.
		ShippingNoticeType string `json:"shipping-notice-type,omitempty"`
		//Вид транспортировки.
		TransportType string `json:"transport-type,omitempty"`
	}

	BatchesResponse struct {
		Batches []Batch `json:"batches"`
		//Список ошибок
		Errors []CreateEntityError `json:"errors"`
		//Список идентификаторов успешно обработанных сущностей
		Ids []int `json:"result-ids"`
	}

	//
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

	CreateEntityError struct {
		//Список кодов ошибок
		Codes []struct {
			Code    string `json:"code"`
			Details string `json:"description"`
		} `json:"error-codes"`
		//Индекс в исходном массиве
		Position int `json:"position"`
	}

	CreateEntityResponse struct {
		//Список ошибок
		Errors []CreateEntityError `json:"errors"`
		//Список идентификаторов успешно обработанных сущностей
		Ids []int `json:"result-ids"`
	}

	DestinationRate struct {
		//Тариф без НДС (коп)
		Rate int `json:"rate"`
		//НДС (коп)
		Vat int `json:"vat"`
	}

	DestinationError struct {
		//код ошибки
		CodeError string `json:"code"`
		//описание ошибки
		DescError string `json:"desc"`
		//под код
		SubCode string `json:"sub-code"`
	}

	DestinationResponse struct {
		//Плата за Авиа-пересылку (коп)
		AviaRate *DestinationRate `json:"avia-rate"`
		//Надбавка за отметку 'Осторожно/Хрупкое'
		FragileRate *DestinationRate `json:"fragile-rate"`
		//Плата за пересылку (коп)
		GroundRate *DestinationRate `json:"ground-rate"`
		//Плата за объявленную ценность (коп)
		InsuranceRate *DestinationRate `json:"insurance-rate"`
		//Надбавка за уведомление о вручении
		NoticeRate *DestinationRate `json:"notice-rate"`
		//Надбавка за негабарит при весе более 10кг
		OversizeRate *DestinationRate `json:"oversize-rate"`
		//Время доставки
		DeliveryTime *DeliveryTime `json:"delivery-time"`
		//Плата всего (коп)
		TotalRate int `json:"total-rate"`
		//Всего НДС (коп)
		TotalVat int `json:"total-vat"`
	}

	CheckInResponse struct {
		//Код ошибки
		ErrorCode string `json:"error-code"`
		//Отослана ли электронная версия ф103/ф103п в ОПС?
		Sent bool `json:"f103-sent"`
	}

	NormalizeAddressRequest struct {
		//Идентификатор записи
		ID string `json:"id"`
		//Оригинальные адрес одной строкой
		OriginalString string `json:"original-address"`
	}

	NormalizeAddress struct {
		//Тип адреса
		AddressType AddressType `json:"address-type"`
		//Район
		Area string `json:"area"`
		//Часть здания: Строение
		Building string `json:"building"`
		//Комментарий:Номер заказа. Внешний идентификатор заказа, который формируется отправителем
		Comment string `json:"comment"`
		//Часть здания: Корпус
		Corpus string `json:"corpus"`
		//Название гостиницы
		Hotel string `json:"hotel"`
		//Часть адреса: Номер здания
		House string `json:"house"`
		//Почтовый индекс
		Index string `json:"index"`
		//Населенный пункт
		Place string `json:"place"`
		//Часть здания: Литера
		Letter string `json:"letter-to"`
		//Область, регион
		Region string `json:"region"`
		//Часть здания: Номер помещения
		Room string `json:"room"`
		//Часть здания: Дробь
		Slash string `json:"slash"`
		//Часть адреса: Улица
		Street string `json:"street-to"`
		//Код качества нормализации адреса
		QualityCode QualityCode `json:"quality-code"`
		//Код проверки нормализации адреса
		ValidationCode ValidationCode `json:"validation-code"`
		//Номер для а/я, войсковая часть, войсковая часть ЮЯ, полевая почта
		NumAddressType string `json:"num-address-type"`
		//Оригинальные адрес одной строкой
		OriginalAddress string `json:"original-address"`
		//Идентификатор записи
		ID string `json:"id"`
	}

	NormalizePhysicalRequest struct {
		//Идентификатор записи
		ID string `json:"id"`
		//Оригинальные адрес одной строкой
		OriginalString string `json:"original-fio"`
	}

	NormalizePhysical struct {
		//Идентификатор записи
		ID string `json:"id"`
		//Оригинальные фамилия, имя, отчество одной строкой
		OriginString string `json:"original-fio"`
		//Отчество
		MiddleName string `json:"middle-name"`
		//Фамилия
		Surname string `json:"surname"`
		//Имя
		Name string `json:"name"`
		//Приемлемое ли качество произведенной очистки?
		Valid bool `json:"valid"`
		//Код качества нормализации адреса
		QualityCode QualityPhysicalCode `json:"quality-code"`
	}

	NormalizePhoneRequest struct {
		//Идентификатор записи
		ID string `json:"id"`
		//Оригинальный номер одной строкой
		OriginalString string `json:"original-phone"`
	}

	NormalizePhone struct {
		//Идентификатор записи
		ID string `json:"id"`
		//Код качества нормализации адреса
		QualityCode QualityPhoneCode `json:"quality-code"`
		//Телефонный номер
		PhoneNumber string `json:"phone-number"`
		//Добавочный номер
		PhoneExtension string `json:"phone-extension"`
		//Код города
		PhoneCityCode string `json:"phone-city-code"`
		//Код страны
		PhoneCountryCode string `json:"phone-country-code"`
		//Оригинальный номер одной строкой
		OriginString string `json:"original-phone"`
	}
)
