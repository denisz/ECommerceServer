package common

type Cursor struct {
	Limit int `json:"limit"`
	Total int `json:"totalElements"`
	Offset int `json:"offset"`
}
