package cdek

type DeliveryMode int

const (
	//дверь-дверь (Д-Д) — Курьер забирает груз у отправителя и доставляет получателю на указанный адрес.
	DeliveryModeDoorDoor DeliveryMode = 1
	//дверь-склад (Д-С) — Курьер забирает груз у отправителя и довозит до склада, получатель забирает груз самостоятельно в ПВЗ (самозабор).
	DeliveryModeDoorWarehouse DeliveryMode = 2
	//склад-дверь (С-Д) — Отправитель доставляет груз самостоятельно до склада, курьер доставляет получателю на указанный адрес.
	DeliveryModeWarehouseDoor DeliveryMode = 3
	//склад-склад (С-С) — Отправитель доставляет груз самостоятельно до склада, получатель забирает груз самостоятельно в ПВЗ (самозабор).
	DeliveryModeWarehouseWarehouse DeliveryMode = 4
)

type DeliveryVariant string

const (
	//курьером
	DeliveryVariantCOURIER DeliveryVariant = "COURIER"
	//ПВЗ СДЭК
	DeliveryVariantPVZCDEK DeliveryVariant = "PVZCDEK"
	//ПВЗ клиента (того кто отправил запрос)
	DeliveryVariantPVZCLIENT DeliveryVariant = "PVZCLIENT"
)

type (

	AuthRequest struct {
		//Дата запроса
		Date	string
		//Идентификатор ИМ, передаваемый СДЭКом.
		Account	string
		//Ключ (см. Протокол обмена)
		Secure	string
	}

	//Размеры
	Dimension struct {
		//Линейная ширина (сантиметры)
		Width int `json:"width"`
		//Линейная высота (сантиметры)
		Height int `json:"height"`
		//Линейная длина (сантиметры)
		Length int `json:"length"`
		//Вес
		Weight float64 `json:"weight"`
	}

	//Тарифы
	Tariff struct {
		//код тарифа
		ID int `json:"id "`
		//заданный приоритет
		Priority int `json:"priority"`
	}

	//Сервис
	ServiceRequest struct {
		//идентификатор номера дополнительной услуги
		ID ServiceCode `json:"id"`
		//параметр дополнительной услуги, если необходимо
		Param int `json:"param"`
	}

	//Сервис
	ServiceResponse struct {
		//идентификатор переданной услуги
		ID ServiceCode `json:"id"`
		//заголовок услуги
		Title string `json:"title"`
		//стоимость услуги без учета НДС
		Price float64 `json:"price"`
	}

	//Запрос на тариф
	DestinationRequest struct {
		//версия используемого API
		Version string `json:"version,omitempty"`
		//планируемая дата оправки заказа в формате “ГГГГ-ММ-ДД”
		DateExecute string `json:"dateExecute,omitempty"`
		//логин
		Account string `json:"authLogin,omitempty"`
		//md5($dateExecute."&".$authPassword)
		Secure string `json:"secure,omitempty"`
		//код города-отправителя в соответствии с кодами городов
		SenderCityID int `json:"senderCityId,omitempty"`
		//код города-отправителя в соответствии с кодами городов
		SenderCityPostCode string `json:"senderCityPostCode,omitempty"`
		//код города-получателя в соответствии с кодами городов
		ReceiverCityID int `json:"receiverCityId,omitempty"`
		//код города-получателя в соответствии с кодами городов
		ReceiverCityPostCode string `json:"receiverCityPostCode,omitempty"`
		//код выбранного тарифа
		TariffID TariffID `json:"tariffId,omitempty"`
		// Список тарифов с приоритетами используется в том случае, если на выбранном направлении у СДЭК может не быть наиболее выгодного для вас, какого-то конкретного тарифа по доставке (например, на направлении Пермь (248) - Астрахань (432) нет тарифа «Посылка склад-склад» (tariffId = 136)). Т. е. тариф  «посылка» действуют не по всем направлениям и не для любого веса груза. В случае задания списка тарифов этот список проверяется на возможность доставки по заданному направлению с заданным весом груза последовательно (начиная с первого с наименьшим приоритетом) и проверка возможности доставки будет проходить до тех пор, пока по очередному тарифу не появится такая возможность. Тогда стоимость будет рассчитана по этому тарифу. В ответе севера будет возвращён
		TariffList []Tariff `json:"tariffList,omitempty"`
		//выбранный режим доставки.
		ModeID DeliveryMode `json:"modeId,omitempty"`
		//Размеры
		Goods []Dimension `json:"goods,omitempty"`
		//Сервисы
		Services []ServiceRequest `json:"services,omitempty"`
	}

	DestinationResponse struct {
		//сумма за доставку в рублях с учетом НДС
		PriceText string `json:"price"`
		//минимальное время доставки в днях
		DeliveryPeriodMin int `json:"deliveryPeriodMin "`
		//максимальное время доставки в днях
		DeliveryPeriodMax int `json:"deliveryPeriodMax "`
		//минимальная дата доставки,  формате 'ГГГГ-ММ-ДД', например “2012-07-29”
		DeliveryDateMin string `json:"deliveryDateMin"`
		//максимальная дата доставки,  формате 'ГГГГ-ММ-ДД', например “2012-07-30”
		DeliveryDateMax string `json:"deliveryDateMax"`
		//код тарифа, по которому посчитана сумма доставки
		TariffID int `json:"tariffId"`
		//ограничение оплаты наличными, появляется только если оно есть
		CashOnDelivery float64 `json:"cashOnDelivery"`
		//ена в валюте с учетом НДС, по которой интернет-магазин работает со СДЭК
		PriceByCurrency int `json:"priceByCurrency"`
		//валюта интернет-магазина (значение из справочника валют см. Приложение 4)
		Currency string `json:"currency"`
		//размер ставки НДС для данного клиента, если у клиента не используется НДС, то данный параметр не будет отображен
		PercentVAT int `json:"percentVAT "`
		//массив с перечнем переданных дополнительных услуг, в котором указывается
		Services []ServiceResponse `json:"services"`
	}

	//Ошибка
	Error struct {
		Code int    `json:"code"`
		Text string `json:"text"`
	}

	//Успешный ответ
	SuccessResponse struct {
		//Результат
		Result DestinationResponse `json:"result"`
	}

	//Неудачный ответ
	FailedResponse struct {
		Error []Error `json:"error"`
	}

	DeliveryRequest struct {
		//Признак международной доставки (1 - международная)
		ForeignDelivery int `xml:"foreignDelivery,attr"`
		//Номер акта приема-передачи/ТТН (сопроводительного документа при передаче груза СДЭК, формируется в системе ИМ), так же используется для удаления заказов
		Number string `xml:"number,attr"`
		//Дата документа (дата заказа) 2015-09-29T00:00:00+06:00
		Date string `xml:"date,attr"`
		//Идентификатор валюты для указания цен (см. Приложение, таблица 6)
		Currency Currency `xml:"currency,attr"`
		//Идентификатор ИМ, передаваемый СДЭКом.
		Account string `xml:"account,attr"`
		//Ключ (см. Протокол обмена)
		Secure string `xml:"secure,attr"`
		//Общее количество заказов в документе
		OrderCount int `xml:"orderCount,attr"`
		//Заказы
		Order []OrderCreateRequest `xml:"order"`
	}

	//Отправление (заказ)
	OrderCreateRequest struct {
		//Номер отправления клиента (должен быть уникален в пределах акта приема-передачи)
		Number string `xml:"number,attr"`
		//Дата инвойса
		DateInvoice string `xml:"dateInvoice,attr"`
		//Код города отправителя из базы СДЭК
		SendCityCode int `xml:"sendCityCode,attr"`
		//Код города получателя из базы СДЭК
		RecCityCode int `xml:"recCityCode,attr"`
		//Почтовый индекс города отправителя
		SendCityPostCode string `xml:"sendCityPostCode,attr"`
		//Почтовый индекс города получателя
		RecCityPostCode string `xml:"recCityPostCode,attr"`
		//Получатель (ФИО)
		RecipientName string `xml:"recipientName,attr"`
		//Email получателя для рассылки уведомлений о движении заказа, для связи в случае недозвона.
		RecipientEmail string `xml:"recipientEmail,attr"`
		//Телефон получателя
		Phone string `xml:"phone,attr"`
		//Код типа тарифа (см. Приложение, таблица 1)
		TariffTypeCode int `xml:"tariffTypeCode,attr"`
		//Комментарий по заказу
		Comment string `xml:"comment,attr"`
		//Истинный продавец. Используется при печати инвойсов для отображения настоящего продавца товара, либо торгового названия
		SellerName string `xml:"sellerName,attr"`
		//Адрес истинного продавца. Используется при печати инвойсов для отображения адреса настоящего продавца товара, либо торгового названия
		SellerAddress string `xml:"sellerAddress,attr"`
		//Грузоотправитель. Используется при печати накладной.
		ShipperName string `xml:"shipperName,attr"`
		//Адрес грузоотправителя. Используется при печати накладной.
		ShipperAddress string `xml:"shipperAddress,attr"`
		//Код валюты наложенного платежа - сумма которую надо взять с получателя (см. Приложение, таблица 6). Если параметр не указан, то считается по значению параметра Currency.
		RecipientCurrency string `xml:"recipientCurrency,attr"`
		//Код валюты объявленной стоимости заказа (всех вложений) (см. Приложение, таблица 6). Если параметр не указан, то считается по значению параметра Currency.
		ItemsCurrency string `xml:"itemsCurrency,attr"`
		//Упаковка
		Package PackageRequest `xml:"package"`
		//Адрес
		Address AddressRequest `xml:"address"`
	}

	//Адрес доставки. В зависимости от режима доставки необходимо указывать либо атрибуты (Street, House, Flat), либо PvzCode
	AddressRequest struct {
		//Улица
		Street string `xml:"street,attr"`
		//Дом, корпус, строение
		House string `xml:"house,attr"`
		//Квартира/Офис
		Flat string `xml:"flat,attr"`
		//Код ПВЗ (см. «Список пунктов выдачи заказов (ПВЗ)»). Атрибут необходим только для заказов с режим доставки «до склада»
		PvzCode string `xml:"pvzCode,attr"`
	}

	//Упаковка (все упаковки передаются в разных тэгах Package)
	PackageRequest struct {
		//Номер упаковки (можно использовать порядковый номер упаковки заказа), уникален в пределах заказа
		Number string `xml:"number,attr"`
		//Штрих-код упаковки (если есть, иначе передавать значение номера упаковки Packege.Number. Параметр используется для оперирования грузом на складах СДЭК), уникален в пределах заказа
		BarCode string `xml:"barCode,attr"`
		//Общий вес (в граммах)
		Weight int `xml:"weight,attr"`
		//Габариты упаковки. Длина (в сантиметрах)
		Length int `xml:"sizeA,attr"`
		//Габариты упаковки. Ширина (в сантиметрах)
		Width int `xml:"sizeB,attr"`
		//Габариты упаковки. Высота (в сантиметрах)
		Height int `xml:"sizeC,attr"`
		//Вложения
		Items []ItemRequest `xml:"item"`
	}

	//Вложение (товар)
	ItemRequest struct {
		//Идентификатор/артикул товара (Уникален в пределах упаковки Package).
		WareKey string `xml:"wareKey,attr"`
		//Объявленная стоимость товара (за единицу товара в указанной валюте, значение >=0).
		CostEx float64 `xml:"costEx,attr"`
		//Объявленная стоимость товара (за единицу товара в рублях, значение >=0) в рублях. С данного значения рассчитывается страховка.
		Cost float64 `xml:"cost,attr"`
		//Оплата за товар при получении (за единицу товара в указанной валюте, значение >=0) — наложенный платеж, в случае предоплаты значение = 0.
		PaymentEx float64 `xml:"paymentEx,attr"`
		//Оплата за товар при получении (за единицу товара в рублях, значение >=0) — наложенный платеж, в случае предоплаты значение = 0.
		Payment float64 `xml:"payment,attr"`
		//Вес нетто (за единицу товара, в граммах)
		Weight int `xml:"weight,attr"`
		//Вес брутто (за единицу товара, в граммах)
		WeightBrutto int `xml:"weightBrutto,attr"`
		//Количество единиц товара (в штуках)
		Amount       int `xml:"amount,attr"`
		//Наименование товара  на английском (может также содержать описание товара: размер, цвет)
		CommentEx string `xml:"commentEx,attr"`
		//Наименование товара на русском (может также содержать описание товара: размер, цвет)
		Comment string `xml:"comment,attr"`
		//Ссылка на сайт интернет-магазина с описанием товара
		Link string `xml:"link,attr"`
	}


	DeleteRequest struct {
		//Номер акта приема-передачи
		Number string `xml:"number,attr"`
		//Дата документа (дата заказа)
		Date string `xml:"date,attr"`
		//Идентификатор ИМ, передаваемый СДЭКом.
		Account	string `xml:"account,attr"`
		//Ключ (см. Протокол обмена)
		Secure	string `xml:"secure,attr"`
		//Общее количество заказов для удаления в документе
		OrderCount	int `xml:"orderCount,attr"`
		//Заказы
		Orders []OrderDeleteRequest `xml:"Order"`
	}

	OrderDeleteRequest struct {
		//Номер отправления клиента
		Number string `xml:"number, attr"`
	}

	//Статус заказа
	StatusReport struct {
		//Дата документа (дата заказа)
		Date string `xml:"date,attr"`
		//Идентификатор ИМ, передаваемый СДЭКом.
		Account	string `xml:"account,attr"`
		//Ключ (см. Протокол обмена)
		Secure	string `xml:"secure,attr"`
		//Атрибут, указывающий на необходимость загружать историю заказов (1-да, 0-нет)
		ShowHistory	int `xml:"showHistory, attr"`
		//Атрибут, указывающий на необходимость загружать список возвратных заказов (1-да, 0-нет)
		ShowReturnOrder 	int `xml:"showReturnOrder, attr"`
		//Атрибут, указывающий на необходимость загружать историю возвратных заказов (1-да, 0-нет)
		ShowReturnOrderHistory 	int `xml:"showReturnOrderHistory, attr"`
		//Период, за который произошло изменение  статуса заказа.
		ChangePeriod StatusReportChangePeriod `xml:"ChangePeriod"`
		//Заказ
		Order StatusReportOrder `xml:"Order"`
	}

	//Период
	StatusReportChangePeriod struct {
		//Дата начала запрашиваемого периода
		DateFirst	string `xml:"dateFirst, attr"`
		//Дата окончания запрашиваемого периода
		DateLast	string `xml:"dateLast, attr"`
	}

	StatusReportOrder struct {
		//Номер отправления СДЭК(присваивается при импорте заказов)
		DispatchNumber int `xml:"dispatchNumber, attr"`
		//Номер отправления клиента
		Number	string `xml:"number, attr"`
		//Дата акта приема-передачи, в котором был передан заказ (2006-01-02)
		Date string `xml:"date, attr"`
	}

	//Запрос
	InfoRequest struct {
		//Дата документа (дата заказа)
		Date string `xml:"date,attr"`
		//Идентификатор ИМ, передаваемый СДЭКом.
		Account	string `xml:"account,attr"`
		//Ключ (см. Протокол обмена)
		Secure	string `xml:"secure,attr"`
		//Отправление (заказ)
		Order InfoRequestOrder `xml:"Order"`
		//Период, за который произошло изменение стоимости услуги доставки
		ChangePeriod InfoRequestChangePeriod `xml:"ChangePeriod"`
	}

	//Период
	InfoRequestChangePeriod struct {
		//Дата начала запрашиваемого периода
		DateBeg	string `xml:"dateBeg, attr"`
		//Дата окончания запрашиваемого периода
		DateEnd	string `xml:"dateEnd, attr"`
	}

	//Заказ отчета
	InfoRequestOrder struct {
		//Номер отправления СДЭК(присваивается при импорте заказов)
		DispatchNumber int `xml:"dispatchNumber, attr"`
		//Номер отправления клиента
		Number	string `xml:"number, attr"`
		//Дата акта приема-передачи, в котором был передан заказ (2006-01-02)
		Date string `xml:"date, attr"`
	}

	//Отчет
	InfoReport struct {
		//Отправление (Заказ)
		Order InfoReportOrder `xml:"Order"`
		//Город отправителя
		SendCity InfoReportCity `xml:"SendCity"`
		//Город получателя
		RecCity InfoReportCity `xml:"RecCity"`
		//Дополнительные услуги к заказам
		AddedService InfoReportAddedService `xml:"AddedService"`
	}

	//Отправление (Заказ)
	InfoReportOrder struct {
		//	Номер отправления клиента
		Number string `xml:"number, attr"`
		//Дата, в которую был передан заказ в базу СДЭК (2006-01-02)
		Date string `xml:"date, attr"`
		//Номер отправления СДЭК(присваивается при импорте заказов)
		DispatchNumber int `xml:"dispatchNumber, attr"`
		//Код типа тарифа (см. Приложение, таблица 1)
		TariffTypeCode	int `xml:"tariffTypeCode, attr"`
		//Расчетный вес (в граммах)
		Weight	float64 `xml:"tariffTypeCode, attr"`
		//Стоимость услуги доставки, руб
		DeliverySum	float64 `xml:"deliverySum, attr"`
		//Дата последнего изменения суммы по услуге доставки
		DateLastChange	string `xml:"dateLastChange, attr"`
		//Режим доставки (1: Д-Д, 2: Д-С, 3: С-Д, 4: С-С)
		DeliveryMode DeliveryMode `xml:"deliveryMode, attr"`
		//Код ПВЗ (см. «Список пунктов выдачи заказов (ПВЗ)»).
		PvzCode	string `xml:"pvzCode, attr"`
		//Вариант доставки (COURIER - курьером, PVZCDEK - ПВЗ СДЭК, PVZCLIENT - ПВЗ клиента (того кто отправил запрос)
		DeliveryVariant	DeliveryVariant `xml:"deliveryVariant, attr"`
	}

	InfoReportCity struct {
		//Код города
		Code int `xml:"code, attr"`
		//Почтовый индекс города
		PostCode string `xml:"postCode, attr"`
		//Название города
		Name string `xml:"name, attr"`
	}

	InfoReportAddedService struct {
		//Тип дополнительной услуги (см. Приложение, таблица 5)
		ServiceCode int `xml:"serviceCode, attr"`
		//Сумма услуги, руб
		Sum	float64 `xml:"sum, attr"`
	}

	//Форма заказа
	OrdersPrint struct {
		//	Дата документа (2006-01-02)
		Date string `xml:"date, attr"`
		//Идентификатор ИМ, передаваемый СДЭКом.
		Account	string `xml:"account, attr"`
		//Ключ (см. Протокол обмена)
		Secure	string `xml:"secure, attr"`
		//Общее количество передаваемых в документе заказов
		OrderCount	int `xml:"orderCount, attr"`
		//Число копий одной квитанции на листе. Рекомендовано указывать не менее 2, одна приклеивается на груз, вторая остается у отправителя.
		CopyCount	int `xml:"copyCount, attr"`
		//Заказы
		Orders []OrderPrint `xml:"Order"`
	}


	//Order	Отправление (заказ)
	OrderPrint struct {
		//Номер отправления СДЭК(присваивается при импорте заказов)
		DispatchNumber int `xml:"dispatchNumber, attr"`
		//Номер отправления клиента
		Number	string `xml:"number, attr"`
		//Дата акта приема-передачи, в котором был передан заказ (2006-01-02)
		Date string `xml:"date, attr"`
	}
)
