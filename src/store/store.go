package store

import (
	"log"
	"context"
	"net/http"
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/gob"
	"store/services/api"
)

type Store struct {
	Config *Config
	API *api.API
	DB *storm.DB
	Scheduler *Scheduler
	HTTPHandler http.Handler
}

func NewStore(config *Config) (*Store, error) {
	DB, err := storm.Open(config.DBFile, storm.Codec(gob.Codec))
	if err != nil {
		return nil, err
	}

	API := api.NewAPI(&api.Config{
		DB: DB,
	})

	scheduler := CreateScheduler(API)
	handlers := CreateMapping(API, append([]string{config.MainServerURL}, config.ExtraURLs...))

	s := &Store{
		Config: config,
		API: API,
		DB: DB,
		Scheduler: scheduler,
		HTTPHandler: handlers,
	}

	return s, nil
}

func(p *Store) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.HTTPHandler.ServeHTTP(w, r)
}

func (p *Store) Shutdown(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("DB close")
			p.DB.Close()
			return
		}
	}
}