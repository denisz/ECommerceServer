package russiaPost


type OperationType int

var (
	//Прием
	OperationType1 OperationType = 1
	//Вручение
	OperationType2 OperationType = 2
	//Возврат
	OperationType3 OperationType = 3
	//Досылка почты
	OperationType4 OperationType = 4
	//Невручение
	OperationType5 OperationType = 5
	//Хранение
	OperationType6 OperationType = 6
	//Временное хранение
	OperationType7 OperationType = 7
	//Обработка
	OperationType8 OperationType = 8
	//Импорт международной почты
	OperationType9 OperationType = 9
	//Экспорт международной почты
	OperationType10 OperationType = 10
	//Прием на таможню
	OperationType11 OperationType = 11
	//Неудачная попытка вручения
	OperationType12 OperationType = 12
	//Регистрация отправки
	OperationType13 OperationType = 13
	//Таможенное оформление
	OperationType14 OperationType = 14
	//Передача на временное хранение
	OperationType15 OperationType = 15
	//Уничтожение
	OperationType16 OperationType = 16
	//Оформление в собственность
	OperationType17 OperationType = 17
	//Регистрация утраты
	OperationType18 OperationType = 18
	//Таможенные платежи поступили
	OperationType19 OperationType = 19
	//Регистрация
	OperationType20 OperationType = 20
	//Доставка
	OperationType21 OperationType = 21
	//Недоставка
	OperationType22 OperationType = 22
	//Поступление на временное хранение
	OperationType23 OperationType = 23
	//Продление срока выпуска таможней
	OperationType24 OperationType = 24
	//Вскрытие
	OperationType25 OperationType = 25
	//Отмена
	OperationType26 OperationType = 26
	//Получена электронная регистрация
	OperationType27 OperationType = 27
	//Присвоение идентификатора
	OperationType28 OperationType = 28
	//Регистрация прохождения в ММПО
	OperationType29 OperationType = 29
	//Отправка SRM
	OperationType30 OperationType = 30
	//Обработка перевозчиком
	OperationType31 OperationType = 31
	//Поступление АПО
	OperationType32 OperationType = 32
	//Международная обработка
	OperationType33 OperationType = 33
	//Электронное уведомление загружено
	OperationType34 OperationType = 34
	//Отказ в курьерской доставке
	OperationType35 OperationType = 35
	//Уточнение вида оплаты доставки
	OperationType36 OperationType = 36
	//Предварительное оформление
	OperationType37 OperationType = 37
	//Задержка для уточнений у отправителя
	OperationType38 OperationType = 38
	//Предварительное таможенное декларирование
	OperationType39 OperationType = 39
	//Таможенный контроль
	OperationType40 OperationType = 40
	//Обработка таможенных платежей
	OperationType41 OperationType = 41
	//Вторая неудачная попытка вручения
	OperationType42 OperationType = 42

)