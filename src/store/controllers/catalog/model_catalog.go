package catalog

import "store/controllers/common"

type DiscountType int32

const (
	DiscountTypePercentage DiscountType  = 0
	DiscountTypeFixedAmount DiscountType = 1
	DiscountTypeFreeShipping DiscountType = 2
)

type Discount struct {
	Type DiscountType `json:"type"`
	Amount int32 `json:"amount"`
}

type Collection struct {
	ID int `storm:"id,increment" json:"id"`
	Name string `storm:"index" json:"name"`
	Picture string `json:"picture"`
}

type Product struct {
	ID int `storm:"id,increment" json:"id"`
	Name string `storm:"index" json:"name"`
	Slug string `storm:"index" json:"slug"`
	Desc string `json:"desc"`
	Spec string `json:"spec"`
	Producer string `json:"producer"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	CurrencyID int `json:"currencyID"`
	CollectionID int `json:"collectionID"`
	Pictures []string `json:"pictures"`
	Discount *Discount `json:"discount"`
}

type PageCollections struct {
	Content []Collection `json:"content"`
	common.Cursor
}

type PageProducts struct {
	Content []Product `json:"content"`
	common.Cursor
}

