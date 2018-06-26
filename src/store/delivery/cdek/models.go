package cdek


type ModeDelivery int

const (
	//дверь-дверь (Д-Д) — Курьер забирает груз у отправителя и доставляет получателю на указанный адрес.
	ModeDeliveryDoorDoor ModeDelivery = 1
	//дверь-склад (Д-С) — Курьер забирает груз у отправителя и довозит до склада, получатель забирает груз самостоятельно в ПВЗ (самозабор).
	ModeDeliveryDoorWarehouse ModeDelivery = 2
	//склад-дверь (С-Д) — Отправитель доставляет груз самостоятельно до склада, курьер доставляет получателю на указанный адрес.
	ModeDeliveryWarehouseDoor ModeDelivery = 3
	//склад-склад (С-С) — Отправитель доставляет груз самостоятельно до склада, получатель забирает груз самостоятельно в ПВЗ (самозабор).
	ModeDeliveryWarehouseWarehouse ModeDelivery = 4
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

	ServiceRequest struct {
		//идентификатор номера дополнительной услуги
		ID ServiceCode `json:"id"`
		//параметр дополнительной услуги, если необходимо
		Param int `json:"param"`
	}

	ServiceResponse struct {
		//идентификатор переданной услуги
		ID ServiceCode `json:"id"`
		//заголовок услуги
		Title string `json:"title"`
		//стоимость услуги без учета НДС
		Price float64 `json:"price"`
	}

	DestinationRequest struct {
		//версия используемого API
		Version string `json:"version,omitempty"`
		//планируемая дата оправки заказа в формате “ГГГГ-ММ-ДД”
		DateExecute string `json:"dateExecute,omitempty"`
		//логин
		Account string `json:"authLogin,omitempty"`
		//md5($dateExecute."&".$authPassword)
		Secure string  `json:"secure,omitempty"`
		//код города-отправителя в соответствии с кодами городов
		SenderCityID int `json:"senderCityId,omitempty"`
		//код города-отправителя в соответствии с кодами городов
		SenderCityPostCode string `json:"senderCityPostCode,omitempty"`
		//код города-получателя в соответствии с кодами городов
		ReceiverCityID int `json:"receiverCityId,omitempty"`
		//код города-получателя в соответствии с кодами городов
		ReceiverCityPostCode string `json:"receiverCityPostCode,omitempty"`
		//код выбранного тарифа
		TariffID int `json:"tariffId,omitempty"`
		// Список тарифов с приоритетами используется в том случае, если на выбранном направлении у СДЭК может не быть наиболее выгодного для вас, какого-то конкретного тарифа по доставке (например, на направлении Пермь (248) - Астрахань (432) нет тарифа «Посылка склад-склад» (tariffId = 136)). Т. е. тариф  «посылка» действуют не по всем направлениям и не для любого веса груза. В случае задания списка тарифов этот список проверяется на возможность доставки по заданному направлению с заданным весом груза последовательно (начиная с первого с наименьшим приоритетом) и проверка возможности доставки будет проходить до тех пор, пока по очередному тарифу не появится такая возможность. Тогда стоимость будет рассчитана по этому тарифу. В ответе севера будет возвращён
		TariffList []Tariff `json:"tariffList,omitempty"`
		//выбранный режим доставки.
		ModeID ModeDelivery `json:"modeId,omitempty"`
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

	Error struct {
		Code int `json:"code"`
		Text string `json:"text"`
	}

	SuccessResponse struct {
		//Результат
		Result DestinationResponse `json:"result"`
	}

	FailedResponse struct {
		Error []Error `json:"error"`
	}

	OrderRequest struct {

	}
)
