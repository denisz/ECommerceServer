package models

import (
	"github.com/asdine/storm"
)

const (
	NodeNamedStore = "Store"
	NodeNamedCarts = "Cart"
	NodeNamedOrders = "Order"
	NodeNamedCatalog = "Catalog"
	NodeNamedSettings = "Settings"
)

// Controller handle all base methods
type Controller struct {
	DB *storm.DB
}

func (p *Controller) GetStore() storm.Node {
	return p.DB.From(NodeNamedStore)
}
