package models

import (
	"github.com/asdine/storm"
)

const (
	NodeNamedStore = "Store"
	NodeNamedCarts = "Cart"
	NodeNamedOrders = "Order"
	NodeNamedBatches = "Batch"
	NodeNamedHistory = "History"
	NodeNamedCatalog = "Catalog"
	NodeNamedSettings = "Settings"
	NodeNamedCDEKCity = "CDEKCity"
	NodeNamedRussiaPost = "RussiaPost"
	NodeNamedAccounting = "Accounting"
)

// Controller handle all base methods
type Controller struct {
	DB *storm.DB
}

func (p *Controller) GetStore() storm.Node {
	return p.DB.From(NodeNamedStore)
}
