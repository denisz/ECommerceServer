package store

import (
	"log"
	"context"
	"net/http"
	"github.com/asdine/storm"
	"store/services/api"
	"store/services/loader"
	. "store/models"
)

type Store struct {
	Config *Config

	// API
	Admin    api.ControllerAdmin
	Cart     api.ControllerCart
	Settings api.ControllerSettings
	Order    api.ControllerOrder
	Account  api.ControllerAccount
	Session  api.ControllerSession
	Catalog  api.ControllerCatalog
	Delivery api.ControllerDelivery
	Sales    api.ControllerSales

	//Loader
	Loader loader.ControllerLoader
}

func createShutdown(db *storm.DB) func(ctx context.Context) {
	return func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("DB close")
				db.Close()
				return
			}
		}
	}
}

func NewStore(config *Config) (http.Handler, func(ctx context.Context), error) {
	db, err := storm.Open("store.db")
	if err != nil {
		return nil, nil, err
	}

	s := &Store{
		Config: config,
		Admin: api.ControllerAdmin{
			Controller: Controller{DB: db},
		},
		Cart: api.ControllerCart{
			Controller: Controller{DB: db},
		},
		Account: api.ControllerAccount{
			Controller: Controller{DB: db},
		},
		Settings: api.ControllerSettings{
			Controller: Controller{DB: db},
		},
		Session: api.ControllerSession{
			Controller: Controller{DB: db},
		},
		Delivery: api.ControllerDelivery{
			Controller: Controller{DB: db},
		},
		Order: api.ControllerOrder{
			Controller: Controller{DB: db},
		},
		Catalog: api.ControllerCatalog{
			Controller: Controller{DB: db},
		},
		Sales: api.ControllerSales{
			Controller: Controller{DB: db},
		},
		Loader: loader.ControllerLoader{
			Controller: Controller{DB: db},
			Config: &loader.Config{
				SpreadSheetID: "13Mr_bOjtMmJ8TivMz3Z5nniT0r92ujlk48m4tXFqSJE",
			},
		},
	}

	return createRouter(s), createShutdown(db), nil
}
