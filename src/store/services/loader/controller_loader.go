package loader

import (
	. "store/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

var (
	RangeBannersName = "Banners"
	RangeProductsName = "Products"
	RangeNotationsName = "Notations"
	RangeCollectionsName = "Collections"
)

type ControllerLoader struct {
	Config *Config
	Controller
}

//Загрузка каталога продуктов
func (p *ControllerLoader) CatalogFromGoogle(c *gin.Context) {
	var err error

	var collections []SheetCollection
	err = UnmarshalSpreadsheet(&collections, p.Config.SpreadSheetID, RangeCollectionsName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var products []SheetProduct
	err = UnmarshalSpreadsheet(&products, p.Config.SpreadSheetID, RangeProductsName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var notations []SheetNotation
	err = UnmarshalSpreadsheet(&notations, p.Config.SpreadSheetID, RangeNotationsName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var catalog = p.GetCatalog()

	tx, err := catalog.Begin(true)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
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
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	// Продукты
	for _, sheetData := range products {
		product := CreateProduct(sheetData)
		err = tx.Save(&product)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	// Описания
	for _, sheetData := range notations {
		product := CreateNotation(sheetData)
		err = tx.Save(&product)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{})
}

//Загрузка рекламных баннеров
func (p *ControllerLoader) AdsFromGoogle(c *gin.Context) {
	var err error
	var banners []SheetBanner
	err = UnmarshalSpreadsheet(&banners, p.Config.SpreadSheetID, RangeBannersName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	settings := Settings {}
	for _, sheetData := range banners {
		if sheetData.Active == 1 {
			settings.Banners = append(settings.Banners, CreateBanner(sheetData))
		}
	}
	db := p.GetSettings()

	tx, err := db.Begin(true)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	defer tx.Rollback()

	err = tx.Set("settings", "754-3010", &settings)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{})
}