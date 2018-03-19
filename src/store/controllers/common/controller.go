package common

import (
	"github.com/asdine/storm"
)

// Controller handle all base methods
type Controller struct {
	DB *storm.DB
	Node storm.Node
}

func (p *Controller) GetStoreNode() storm.Node {
	return p.Node
}