package api

import (
	. "store/models"
	"store/services/updater"
)

type API struct {
	Config   *Config
	Admin    ControllerAdmin
	Cart     ControllerCart
	Settings ControllerSettings
	Order    ControllerOrder
	Form     ControllerForms
	Account  ControllerAccount
	Catalog  ControllerCatalog
	Sales    ControllerSales
	Batches  ControllerBatches
	Loader   ControllerUpdater
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
		Batches: ControllerBatches{
			Controller: Controller{DB: db},
		},
		Loader: ControllerUpdater{
			Controller: Controller{DB: db},
			Config: &updater.Config{
				SpreadSheetID: "13Mr_bOjtMmJ8TivMz3Z5nniT0r92ujlk48m4tXFqSJE",
			},
		},
	}
}
