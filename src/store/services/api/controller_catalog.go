package api

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	. "store/models"
	"fmt"
)

type ControllerCatalog struct {
	Controller
}

// Коллекция
func (p *ControllerCatalog) GetCollectionBySKU(sku string) (*Collection, error){
	var collection Collection
	err := p.GetStore().From(NodeNamedCatalog).One("SKU", sku, &collection)

	if err == storm.ErrNotFound {
		return nil, err
	}

	return &collection, nil
}

// Список коллекции
func (p *ControllerCatalog) GetAllCollections() (*PageCollections, error) {
	var collections []Collection
	err := p.GetStore().From(NodeNamedCatalog).All(&collections)
	if err != nil {
		return nil, err
	}

	return &PageCollections{
		Content: collections,
		Cursor: Cursor{
			Total:  len(collections),
			Limit:  len(collections),
			Offset: 0,
		},
	}, nil
}

// Список товаров
func (p *ControllerCatalog) GetProductsByCollectionSKU(sku string, pagination Pagination) (*PageProducts, error) {
	var products []Product
	err := p.GetStore().From(NodeNamedCatalog).
		Find("CollectionSKU", sku, &products, storm.Limit(pagination.Limit), storm.Skip(pagination.Offset))

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().From(NodeNamedCatalog).
		Select(q.Eq("CollectionSKU", sku)).
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
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//Поиск по наименованию товара
func (p *ControllerCatalog) SearchProductsWithFilter(filter FilterCatalog, pagination Pagination) (*PageProducts, error) {
	matcher := q.True()

	if len(filter.CollectionSKU) > 0 {
		matcher = q.And(matcher, q.Eq("CollectionSKU", filter.CollectionSKU))
	}

	if len(filter.Query) > 0 {
		matcher = q.And(matcher, q.Re("Name", fmt.Sprintf("^%s", filter.Query)))
	}

	if len(filter.Producer) > 0 {
		matcher = q.And(matcher, q.Eq("Producer", filter.Producer))
	}

	var products []Product
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
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//Продукт
func (p *ControllerCatalog) GetProductBySKU(sku string) (*Product, error) {
	var product Product
	err := p.GetStore().From(NodeNamedCatalog).One("SKU", sku, &product)

	if err == storm.ErrNotFound {
		return nil, err
	}

	product.PriceCalculate()
	return &product, nil
}

//Описания продукта
func (p *ControllerCatalog) GetNotationBySKU(sku string) (*Notation, error) {
	var notation Notation
	err := p.GetStore().From(NodeNamedCatalog).One("SKU", sku, &notation)

	if err == storm.ErrNotFound {
		return nil, err
	}

	return &notation, nil
}
