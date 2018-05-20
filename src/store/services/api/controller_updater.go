package api

import (
	. "store/models"
	"store/services/updater"
	"fmt"
)

var (
	RangeMediaName = "Media"
	RangePricesName = "Prices"
	RangeBannersName = "Banners"
	RangeProductsName = "Products"
	RangeNotationsName = "Notations"
	RangeCollectionsName = "Collections"
)

type ControllerUpdater struct {
	Config *updater.Config
	Controller
}

//Загрузка каталога продуктов
func (p *ControllerUpdater) CatalogFromGoogle() error {
	var err error

	var collections []updater.SheetCollection
	err = updater.UnmarshalSpreadsheet(&collections, p.Config.SpreadSheetID, RangeCollectionsName)
	if err != nil {
		return err
	}

	var products []updater.SheetProduct
	err = updater.UnmarshalSpreadsheet(&products, p.Config.SpreadSheetID, RangeProductsName)
	if err != nil {
		return err
	}

	var notations []updater.SheetNotation
	err = updater.UnmarshalSpreadsheet(&notations, p.Config.SpreadSheetID, RangeNotationsName)
	if err != nil {
		return err
	}

	var media []updater.SheetProductMedia
	err = updater.UnmarshalSpreadsheet(&media, p.Config.SpreadSheetID, RangeMediaName)
	if err != nil {
		return err
	}

	var catalog = p.GetStore().From(NodeNamedCatalog)

	tx, err := catalog.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//remove all products
	err = tx.Drop(&Product{})
	if err != nil {
		fmt.Printf("Drop error: %v \n", err)
	}
	//remove all collections
	err = tx.Drop(&Collection{})
	if err != nil {
		fmt.Printf("Drop error: %v \n", err)
	}

	//remove all notations
	err = tx.Drop(&Notation{})
	if err != nil {
		fmt.Printf("Drop error: %v \n", err)
	}

	tx.ReIndex(Product{})
	tx.ReIndex(Notation{})
	tx.ReIndex(Collection{})

	// Категории
	for _, sheetData := range collections {
		collection := updater.CreateCollection(sheetData)
		err = tx.Save(&collection)
		if err != nil {
			return err
		}
	}

	// Продукты
	for _, sheetData := range products {
		if sheetData.Disabled { continue }

		for _, m := range media {
			if m.SKU == sheetData.SKU {
				sheetData.Pictures = m.Pictures
			}
		}

		product := updater.CreateProduct(sheetData)
		err = tx.Save(&product)
		if err != nil {
			return err
		}
	}

	// Описания
	for _, sheetData := range notations {
		product := updater.CreateNotation(sheetData)
		err = tx.Save(&product)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

//Загрузка рекламных баннеров
func (p *ControllerUpdater) AdsFromGoogle() error {
	var err error
	var banners []updater.SheetBanner
	err = updater.UnmarshalSpreadsheet(&banners, p.Config.SpreadSheetID, RangeBannersName)
	if err != nil {
		return err
	}

	settings := Settings {}
	for _, sheetData := range banners {
		if sheetData.Active {
			settings.Banners = append(settings.Banners, updater.CreateBanner(sheetData))
		}
	}
	db := p.GetStore().From(NodeNamedSettings)

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.Set("settings", "754-3010", &settings)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

//Обновление цен из таблиц
func (p *ControllerUpdater) PriceFromGoogle() error {
	var err error

	var prices []updater.SheetPrice
	err = updater.UnmarshalSpreadsheet(&prices, p.Config.SpreadSheetID, RangePricesName)
	if err != nil {
		return err
	}

	//магазин
	store := p.GetStore()
	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//каталог
	catalog := tx.From(NodeNamedCatalog)

	//все продукты
	var products []Product
	catalog.AllByIndex("ID", &products)

	//обновляем цены на всех товарах
	for _, product := range products {
		for _, price := range prices {
			if price.SKU == product.SKU {
				product.Price = Price(price.Price * 100)
				err := catalog.Save(&product)
				if err != nil {
					return err
				}
			}
		}
	}

	//фиксируем транзакцию
	tx.Commit()

	return nil
}