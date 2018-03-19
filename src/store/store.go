package store

import (
	"log"
	"context"
	"net/http"
	"github.com/asdine/storm"
	"store/controllers/cart"
	"store/controllers/account"
	"store/controllers/info"
	"store/controllers/admin"
	"store/controllers/session"
	"store/controllers/common"
	"store/controllers/shipment"
	"store/controllers/order"
	"store/controllers/catalog"
	"store/services/updater"
	"fmt"
)

type Store struct {
	Config   *Config
	Admin    admin.Controller
	Cart     cart.Controller
	Info     info.Controller
	Order    order.Controller
	Account  account.Controller
	Session  session.Controller
	Catalog  catalog.Controller
	Shipment shipment.Controller
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
	node := db.From("store")

	s := &Store{
		Config: config,
		Admin: admin.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Cart: cart.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Account: account.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Info: info.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Session: session.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Shipment: shipment.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Order: order.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
		Catalog: catalog.Controller{
			Controller: common.Controller{DB: db, Node: node},
		},
	}

	err = updater.Updater(node, &updater.Config {
		SpreadSheetID: "13Mr_bOjtMmJ8TivMz3Z5nniT0r92ujlk48m4tXFqSJE",
	})

	if err != nil {
		fmt.Print(err)
	}

	return createRouter(s), createShutdown(db), nil
}


