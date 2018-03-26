package models

import (
	"github.com/asdine/storm"
)

// Controller handle all base methods
type Controller struct {
	DB   *storm.DB
}

func (p *Controller) GetStore() storm.Node {
	return p.DB.From("Store")
}

func (p *Controller) GetCatalog() storm.Node {
	return p.GetStore().From("Catalog")
}

func (p *Controller) GetSettings() storm.Node {
	return p.GetStore().From("Settings")
}

