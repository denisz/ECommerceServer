package cdek


import "errors"

var (
	//0
	ErrInterServerError = errors.New("Внутренняя ошибка на сервере. Обратитесь к программистам компании СДЭК для исправления")
	//1
	ErrNoValidVersionAPI = errors.New("Указанная вами версия API не поддерживается")
	//2
	ErrAuthentication = errors.New("Ошибка авторизации")
	//3
	ErrDeliveryImpossible = errors.New("Невозможно осуществить доставку по этому направлению при заданных условиях")
	//4
	ErrNoValidPlace = errors.New("Ошибка при указании параметров места")
	//5
	ErrEmptyPlace = errors.New("Не задано ни одного места для отправления")
	//6
	ErrEmptyTariff = errors.New("Не задан тариф или список тарифов")
	//7
	ErrEmptySendCity = errors.New("Не задан город-отправитель")
	//8
	ErrEmptyRecCity = errors.New("задан город-получатель")
	//9
	ErrEmptyDate = errors.New("При авторизации не задана дата планируемой отправки")
	//10
	ErrNoValidModeID = errors.New("Ошибка задания режима доставки")
	//11
	ErrNoSupportFormat = errors.New("Неправильно задан формат данных")
	//12
	ErrDecodeFormat = errors.New("Ошибка декодирования данных. Ожидается <json или jsop>")
	//13
	 ErrEmptyPostCode = errors.New("Почтовый индекс города-отправителя отсутствует в базе СДЭК")
	//14
	ErrImplicitSendCityCode = errors.New("Невозможно однозначно идентифицировать город-отправитель по почтовому индексу")
	//15
	ErrNotFoundsPostCode = errors.New("Почтовый индекс города-получателя отсутствует в базе СДЭК")
	//16
	ErrImplicitRecCityCode = errors.New("Невозможно однозначно идентифицировать город-получатель по почтовому индексу")
)
