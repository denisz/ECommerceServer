package models


//Время доставки для почты России
type RussiaPostDeliveryPeriod struct {
	// Индентификатор
	ID int `storm:"id,increment" json:"id"`
	Region string `json:"region"`
	Capital string `json:"capital"`
	EMC DeliveryPeriod `json:"emc"`
	Rapid DeliveryPeriod `json:"rapid"`
}
