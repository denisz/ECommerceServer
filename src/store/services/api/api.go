package api

import (
	. "store/models"
)

type API struct {
	Config     *Config
	Admin      ControllerAdmin
	Cart       ControllerCart
	Settings   ControllerSettings
	Order      ControllerOrder
	Form       ControllerForms
	Account    ControllerAccount
	Catalog    ControllerCatalog
	Sales      ControllerSales
	Batch      ControllerBatch
	Loader     ControllerUpdater
	Accounting ControllerAccounting
}

func NewAPI(config *Config) *API {
	db := config.DB

	return &API{
		Admin: ControllerAdmin{
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
		Form: ControllerForms{
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
		Batch: ControllerBatch{
			Controller: Controller{DB: db},
		},
		Loader: ControllerUpdater{
			Controller:    Controller{DB: db},
			SpreadSheetID: "13Mr_bOjtMmJ8TivMz3Z5nniT0r92ujlk48m4tXFqSJE",
		},
		Accounting: ControllerAccounting{
			Controller: Controller{DB: db},
			FileID:     "1br5HTFWzYWTrALqfBb3SYWbFZOzhe8lj",
		},
	}
}
