package order

import (
	"time"
	"store/controllers/account"
	"store/controllers/catalog"
)

type Status int32

const (
	OrderStatusDraft 	  Status = 0 // новый заказ
	OrderStatusPending    Status = 1 // формированный заказ
	OrderStatusProcessing Status = 2 // в обработке
	OrderStatusClosed     Status = 3 // закрыт
	OrderStatusCanceled   Status = 4 // отменен
)

type Receipt struct {
	ID int `storm:"id,increment"`
	OrderID int
	Response string
	Provider string
}

type Shipping struct {
	Code string `json:"code"`
	Provider string `json:"provider"`
}

type Item struct {
	Product catalog.Product `json:"productID"`
	Amount int `json:"amount"`
}

type Order struct {
	ID            int              `storm:"id,increment" json:"id"`
	Items         []Item           `json:"items"`
	Address       account.Address  `json:"address"`
	Receipt       Receipt          `json:"-"`
	Shipping      Shipping         `json:"shipping"`
	Discount      catalog.Discount `json:"discount"`
	Status        Status           `json:"status"`
	UserID        int              `json:"userID"`
	Invoice       int              `json:"invoice"`
	TaxPrice      int              `json:"taxPrice"`
	TotalPrice    int              `json:"totalPrice"`
	ShippingPrice int              `json:"shippingPrice"`
	Comment       string           `json:"comment"`
	CreatedAt     time.Time        `json:"-"`
}

type History struct {
	ID int `storm:"id,increment"`
	OrderID int
	OperatorID int
	Comment string
	Status string
}

