package cmd

import (
	"fmt"
	"store/controllers/catalog"
	"github.com/asdine/storm"
	"log"
)

func main() {
	db, err := storm.Open("store.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.Drop("store")
	if err != nil {
		log.Fatal(err)
	}

	c := db.From("store")

	for i := 1; i <= 6; i++ {
		collection := &catalog.Collection{
			Name: fmt.Sprintf("Collection %d", i),
		}

		c.Save(collection)

		for i := 1; i <= 10; i++ {
			product := &catalog.Product{
				Quantity: 10,
				CollectionID: collection.ID,
				Slug: fmt.Sprintf("product_%d", i),
				Name: fmt.Sprintf("Product %d", i),
				Price: 3500,
				CurrencyID: 643,
				Pictures: []string {
					fmt.Sprintf("/products/%d/main.jpg", i),
				},
				Producer: fmt.Sprintf("Producer %d", i),
				Desc: fmt.Sprintf("Description %d", i),
			}

			if i == 1 {
				product.Discount = &catalog.Discount{
					Type: catalog.DiscountTypePercentage,
					Amount: 10,
				}
			}

			c.Save(product)
		}
	}

	var collections []catalog.Collection
	err = c.All(&collections)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" %v", collections)

	var products []catalog.Product
	err = c.All(&products)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" %v", products)
}
