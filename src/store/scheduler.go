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

func (p *Scheduler) SearchReceivedReports() {
	err := p.API.Accounting.CheckDeliveryReports()
	if err != nil {
		log.Error().Err(err)
	}

	err = p.API.Accounting.UpdateQuantity()
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

	gocron.Every(1).
		Day().
		At("23:10").
		Do(scheduler.SearchReceivedReports)


	gocron.Start()

	return scheduler
}
