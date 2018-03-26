package models

type (
	//Настройки сайта
	Settings struct {
		Banners []Banner `json:"banners"`
	}

	//Баннер
	Banner struct {
		Image string `json:"img"`
		Href string `json:"href"`
	}
)