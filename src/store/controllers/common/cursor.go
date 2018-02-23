package common

type Cursor struct {
	Last bool `json:"last"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}
