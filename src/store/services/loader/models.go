package loader

import (
	. "store/models"
	"fmt"
)

type (
	SheetProduct struct {
		Quantity      int      `sheet:"Количество"`
		Price         int      `sheet:"Цена"`
		Weight        int      `sheet:"Вес"`
		Name          string   `sheet:"Наименование"`
		CollectionSKU string   `sheet:"Категория"`
		Producer      string   `sheet:"Производитель"`
		Factor        string   `sheet:"Фасовка"`
		Form          string   `sheet:"Форма"`
		SKU           string   `sheet:"Артикул"`
		Discount      string   `sheet:"Скидка"`
		Pictures      []string `sheet:"Картинки"`
	}

	SheetCollection struct {
		SKU  string `sheet:"Артикул"`
		Name string `sheet:"Наименование"`
	}

	SheetNotation struct {
	}

	SheetBanner struct {
		Image  string `sheet:"Изображение"`
		Active int    `sheet:"Активен"`
		Href   string `sheet:"Переход"`
	}
)

func CreateBanner(sheetData SheetBanner) Banner {
	banner := Banner{
		Image: sheetData.Image,
		Href:  sheetData.Href,
	}

	return banner
}

func CreateCollection(sheetData SheetCollection) Collection {
	return Collection{
		Name: sheetData.Name,
		SKU:  sheetData.SKU,
	}
}

func CreateProduct(sheetData SheetProduct) Product {
	product := Product{
		Name:          sheetData.Name,
		Producer:      sheetData.Producer,
		Form:          sheetData.Form,
		Factor:        sheetData.Factor,
		Weight:        sheetData.Weight,
		SKU:           sheetData.SKU,
		Quantity:      sheetData.Quantity,
		CollectionSKU: sheetData.CollectionSKU,
		Price:         sheetData.Price * 100, // 100 копеек
		Pictures:      []string{},
	}

	for _, p := range sheetData.Pictures {
		product.Pictures = append(product.Pictures, fmt.Sprintf("products/%s/%s.jpg", product.SKU, p))
	}

	if len(sheetData.Discount) > 0 {
		if num, percent, ok := percent(sheetData.Discount); ok && num > 0 {
			if percent {
				product.Discount = &Discount{
					Type:   DiscountTypePercentage,
					Amount: num,
				}
			} else {
				product.Discount = &Discount{
					Type:   DiscountTypeFixedAmount,
					Amount: num,
				}
			}
		}
	}

	return product
}
