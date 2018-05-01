package store

import (
	"github.com/jasonlvhit/gocron"
	"store/services/api"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	API *api.API
}


func (p *Scheduler) ClearExpiredOrders() {
	err := p.API.Order.ClearExpiredOrders()

	if err != nil {
		log.Error().Err(err)
	}
}

func CreateScheduler(API *api.API) *Scheduler {
	scheduler := &Scheduler{
		API: API,
	}

	gocron.Every(1).
		Day().
		At("23:00").
		Do(scheduler.ClearExpiredOrders)


	gocron.Start()

	return scheduler
}
