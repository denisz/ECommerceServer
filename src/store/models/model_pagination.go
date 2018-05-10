package models


type Pagination struct {
	// Лимит
	Limit int `form:"limit" json:"limit"`
	// Смещение
	Offset int `form:"offset" json:"offset"`
	// Страница
	Page int `form:"page" json:"page"`
}
