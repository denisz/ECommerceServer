package models

type BannerType string

const (
	BannerTypeMain BannerType = "main"
	BannerTypeBrand BannerType = "brand"
)

type (
	//Настройки сайта
	Settings struct {
		Banners []Banner `json:"banners"`
	}

	//Баннер
	Banner struct {
		Image string `json:"img"`
		Href string `json:"href"`
		Active bool `json:"active"`
		Type BannerType `json:"type"`
	}
)