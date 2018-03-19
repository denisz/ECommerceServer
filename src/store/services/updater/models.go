package updater

import (
	"store/controllers/catalog"
	"fmt"
)

type (
	SheetProduct struct {
		ID         int      `sheet:"Номер"`
		Quantity   int      `sheet:"Количество"`
		Price      int      `sheet:"Цена"`
		Weight     int      `sheet:"Вес"`
		Name       string   `sheet:"Наименование"`
		Collection string   `sheet:"Категория"`
		Producer   string   `sheet:"Производитель"`
		Factor     string   `sheet:"Фасовка"`
		Form       string   `sheet:"Форма"`
		SKU        string   `sheet:"Артикул"`
		Discount   string   `sheet:"Скидка"`
		Pictures   []string `sheet:"Картинки"`
	}

	SheetCollection struct {
		ID   int    `sheet:"Номер"`
		Name string `sheet:"Наименование"`
	}

	SheetNotation struct {
	}
)

func CreateCollection(sheetData SheetCollection) catalog.Collection {
	return catalog.Collection{
		ID:   sheetData.ID,
		Name: sheetData.Name,
	}
}

func CreateProduct(sheetData SheetProduct) catalog.Product {
	product := catalog.Product{
		ID:       sheetData.ID,
		Name:     sheetData.Name,
		Producer: sheetData.Producer,
		Form:     sheetData.Form,
		Factor:   sheetData.Factor,
		Weight:   sheetData.Weight,
		SKU:      sheetData.SKU,
		Quantity: sheetData.Quantity,
		Price:    sheetData.Price,
		Pictures: []string{},
	}

	for _, p := range sheetData.Pictures {
		product.Pictures = append(product.Pictures, fmt.Sprintf("products/%s/%s.jpg", underscoreString(sheetData.SKU), p))
	}

	if len(sheetData.Discount) > 0 {
		if num, percent, ok := percent(sheetData.Discount); ok && num > 0 {
			if percent {
				product.Discount = &catalog.Discount{
					Type:   catalog.DiscountTypePercentage,
					Amount: num,
				}
			} else {
				product.Discount = &catalog.Discount{
					Type:   catalog.DiscountTypeFixedAmount,
					Amount: num,
				}
			}
		}
	}

	return product
}
