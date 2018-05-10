package models

type Cursor struct {
	// Лимит
	Limit int `json:"limit"`
	// Общее количество
	Total int `json:"totalElements"`
	// Смещение
	Offset int `json:"offset"`
}
