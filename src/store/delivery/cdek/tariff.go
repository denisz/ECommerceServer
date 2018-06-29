package cdek


type TariffID int

const (
	//Посылка дверь-дверь	дверь-дверь (Д-Д)
	TariffIDDoorDoor TariffID = 139
	//Посылка дверь-склад	склад-дверь (Д-С)
	TariffIDDoorWarehouse TariffID = 138
	//Посылка склад-дверь	склад-дверь (С-Д)
	TariffIDWarehouseDoor TariffID = 137
	//Посылка склад-склад	склад-склад (С-С)	до 30 кг.	Посылка	Услуга экономичной доставки товаров по России для компаний, осуществляющих дистанционную торговлю.
	TariffIDWarehouseWarehouse TariffID = 136
)
