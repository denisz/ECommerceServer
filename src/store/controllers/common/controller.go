package common

import (
	"github.com/asdine/storm"
)

// Controller handle all base methods
type Controller struct {
	DB *storm.DB
	StoreNode storm.Node
}

func (p *Controller) GetStoreNode() storm.Node {
	return p.DB.From("store")
}