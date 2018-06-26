package models


type CDEKCity struct {
	// Индентификатор
	ID int `storm:"id,increment" json:"id"`
	//Код в системе CDEK
	Code int `storm:"index" json:"code"`
	// Название
	Name string `json:"name"`
	// Область, район
	District string `json:"district"`
	//Postcode
	PostCode string `storm:"index" json:"postcode"`
}
