package api

import (
	. "store/models"
	"store/services/loader"
	"fmt"
)

var (
	RangeMediaName = "Media"
	RangeBannersName = "Banners"
	RangeProductsName = "Products"
	RangeNotationsName = "Notations"
	RangeCollectionsName = "Collections"
)

type ControllerLoader struct {
	Config *loader.Config
	Controller
}

//Загрузка каталога продуктов
func (p *ControllerLoader) CatalogFromGoogle() error {
	var err error

	var collections []loader.SheetCollection
	err = loader.UnmarshalSpreadsheet(&collections, p.Config.SpreadSheetID, RangeCollectionsName)
	if err != nil {
		return err
	}

	var products []loader.SheetProduct
	err = loader.UnmarshalSpreadsheet(&products, p.Config.SpreadSheetID, RangeProductsName)
	if err != nil {
		return err
	}

	var notations []loader.SheetNotation
	err = loader.UnmarshalSpreadsheet(&notations, p.Config.SpreadSheetID, RangeNotationsName)
	if err != nil {
		return err
	}

	var media []loader.SheetProductMedia
	err = loader.UnmarshalSpreadsheet(&media, p.Config.SpreadSheetID, RangeMediaName)
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
		collection := loader.CreateCollection(sheetData)
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

		product := loader.CreateProduct(sheetData)
		err = tx.Save(&product)
		if err != nil {
			return err
		}
	}

	// Описания
	for _, sheetData := range notations {
		product := loader.CreateNotation(sheetData)
		err = tx.Save(&product)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

//Загрузка рекламных баннеров
func (p *ControllerLoader) AdsFromGoogle() error {
	var err error
	var banners []loader.SheetBanner
	err = loader.UnmarshalSpreadsheet(&banners, p.Config.SpreadSheetID, RangeBannersName)
	if err != nil {
		return err
	}

	settings := Settings {}
	for _, sheetData := range banners {
		if sheetData.Active {
			settings.Banners = append(settings.Banners, loader.CreateBanner(sheetData))
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