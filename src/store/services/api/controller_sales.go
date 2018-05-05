package api

import (
	. "store/models"
	"github.com/asdine/storm/q"
	"github.com/asdine/storm"
)

type ControllerSales struct {
	Controller
}


func(p *ControllerSales) GetProducts(pagination Pagination) (*PageProducts, error) {
	var products []Product

	matcher := q.Not(q.Eq("Discount", nil))

	err := p.GetStore().From(NodeNamedCatalog).
		Select(matcher).
		Limit(pagination.Limit).
		Skip(pagination.Offset).
		Find(&products)

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().From(NodeNamedCatalog).
		Select(matcher).
		Count(new(Product))

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		product.PriceCalculate()
	}

	return &PageProducts{
		Content: products,
		Cursor: Cursor{
			Total: total,
			Limit: pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

func (p *ControllerSales) UpdateProducts() error {
	return nil
}