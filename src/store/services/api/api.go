package api

import (
	. "store/models"
	"store/services/loader"
)

type API struct {
	Config   *Config
	Admin    ControllerAdmin
	Cart     ControllerCart
	Settings ControllerSettings
	Order    ControllerOrder
	Account  ControllerAccount
	Catalog  ControllerCatalog
	Sales    ControllerSales
	Loader   ControllerLoader
}

func NewAPI(config *Config) *API {
	db := config.DB

	return &API{
		Admin: ControllerAdmin {
			Controller: Controller{DB: db},
		},
		Cart: ControllerCart{
			Controller: Controller{DB: db},
		},
		Settings: ControllerSettings{
			Controller: Controller{DB: db},
		},
		Order: ControllerOrder{
			Controller: Controller{DB: db},
		},
		Account: ControllerAccount{
			Controller: Controller{DB: db},
		},
		Catalog: ControllerCatalog{
			Controller: Controller{DB: db},
		},
		Sales: ControllerSales{
			Controller: Controller{DB: db},
		},
		Loader: ControllerLoader{
			Controller: Controller{DB: db},
			Config: &loader.Config{
				SpreadSheetID: "13Mr_bOjtMmJ8TivMz3Z5nniT0r92ujlk48m4tXFqSJE",
			},
		},
	}
}
