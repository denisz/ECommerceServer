package updater

import (
	"github.com/asdine/storm"
	"store/controllers/catalog"
	"fmt"
)

var (
	RangeCollectionsName = "Collections"
	RangeProductsName = "Products"
)

func Updater(db storm.Node, config *Config) error {
	var err error
	var collections []SheetCollection
	err = UnmarshalSpreadsheet(&collections, config.SpreadSheetID, RangeCollectionsName)
	if err != nil {
		return err
	}

	var products []SheetProduct
	err = UnmarshalSpreadsheet(&products, config.SpreadSheetID, RangeProductsName)
	if err != nil {
		return err
	}

	//remove all products
	db.Drop(&catalog.Product{})
	//remove all collections
	db.Drop(&catalog.Collection{})

	for _, dataSheet := range collections {
		collection := CreateCollection(dataSheet)
		err = db.Save(&collection)
		if err != nil {
			fmt.Print(err)
		}
	}

	for _, dataSheet := range products {
		product := CreateProduct(dataSheet)
		for _, collection := range collections {
			if collection.Name == dataSheet.Collection {
				product.CollectionID = collection.ID
				break
			}
		}
		err = db.Save(&product)
		if err != nil {
			fmt.Print(err)
		}
	}


	return nil
}
