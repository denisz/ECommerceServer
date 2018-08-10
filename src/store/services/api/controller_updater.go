package api

import (
	. "store/models"
	"store/services/gdrv"
	"fmt"
)

var (
	RangeMediaName        = "Media"
	RangePricesName       = "Prices"
	RangeBannersName      = "Banners"
	RangeProductsName     = "Products"
	RangeCDEKCityName     = "CDEKCity"
	RangeNotationsName    = "Notations"
	RangeRussiaPostPeriod = "RussiaPost"
	RangeCollectionsName  = "Collections"
)

type ControllerUpdater struct {
	SpreadSheetID string
	Controller
}

//Загрузка каталога продуктов
func (p *ControllerUpdater) CatalogFromGoogle() error {
	var err error

	var collections []SheetCollection
	err = gdrv.UnmarshalSpreadsheet(&collections, p.SpreadSheetID, RangeCollectionsName)
	if err != nil {
		return err
	}

	var products []SheetProduct
	err = gdrv.UnmarshalSpreadsheet(&products, p.SpreadSheetID, RangeProductsName)
	if err != nil {
		return err
	}

	var notations []SheetNotation
	err = gdrv.UnmarshalSpreadsheet(&notations, p.SpreadSheetID, RangeNotationsName)
	if err != nil {
		return err
	}

	var media []SheetProductMedia
	err = gdrv.UnmarshalSpreadsheet(&media, p.SpreadSheetID, RangeMediaName)
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
		collection := CreateCollection(sheetData)
		err = tx.Save(&collection)
		if err != nil {
			return err
		}
	}

	// Продукты
	for _, sheetData := range products {
		if sheetData.Disabled {
			continue
		}

		for _, m := range media {
			if m.SKU == sheetData.SKU {
				sheetData.Pictures = m.Pictures
			}
		}

		product := CreateProduct(sheetData)
		err = tx.Save(&product)
		if err != nil {
			return err
		}
	}

	// Описания
	for _, sheetData := range notations {
		product := CreateNotation(sheetData)
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
	var banners []SheetBanner
	err = gdrv.UnmarshalSpreadsheet(&banners, p.SpreadSheetID, RangeBannersName)
	if err != nil {
		return err
	}

	settings := Settings{}
	for _, sheetData := range banners {
		if sheetData.Active {
			settings.Banners = append(settings.Banners, CreateBanner(sheetData))
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

	var prices []SheetPrice
	err = gdrv.UnmarshalSpreadsheet(&prices, p.SpreadSheetID, RangePricesName)
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

//Обновления списка городов
func (p *ControllerUpdater) CDEKCityFromGoogle() error {
	var err error

	var collections []SheetCDEKCity
	err = gdrv.UnmarshalSpreadsheet(&collections, p.SpreadSheetID, RangeCDEKCityName)
	if err != nil {
		return err
	}

	catalog := p.DB.From(NodeNamedCDEKCity)
	tx, err := catalog.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//remove all CDEKCity
	err = tx.Drop(&CDEKCity{})
	if err != nil {
		fmt.Printf("Drop error: %v \n", err)
	}

	tx.ReIndex(CDEKCity{})

	// Описания
	for _, sheetData := range collections {
		for _, postcode := range sheetData.PostCodeList {
			city := CreateCDEKCity(sheetData, postcode)
			err = tx.Save(&city)
			if err != nil {
				return err
			}
		}
	}

	tx.Commit()
	return nil
}

//Обновления списка городов
func (p *ControllerUpdater) RussiaPostFromGoogle() error {
	var err error

	var collections []SheetRussiaPost
	err = gdrv.UnmarshalSpreadsheet(&collections, p.SpreadSheetID, RangeRussiaPostPeriod)
	if err != nil {
		return err
	}

	catalog := p.DB.From(NodeNamedRussiaPost)
	tx, err := catalog.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//remove all RussiaPostDeliveryPeriod
	err = tx.Drop(&RussiaPostDeliveryPeriod{})
	if err != nil {
		fmt.Printf("Drop error: %v \n", err)
	}

	tx.ReIndex(RussiaPostDeliveryPeriod{})

	for _, sheetData := range collections {
		if len(sheetData.DeliveryTimeRapid) > 0 && len(sheetData.DeliveryTimeEMC) > 0 {
			time := CreateRussiaPostDeliveryPeriod(sheetData)
			err = tx.Save(&time)
			if err != nil {
				return err
			}
		}
	}

	tx.Commit()
	return nil
}
