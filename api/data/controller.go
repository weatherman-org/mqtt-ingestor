package data

import (
	"github.com/go-chi/chi/v5"
	db "github.com/weatherman-org/telemetry/db/sqlc"
	"github.com/weatherman-org/telemetry/util"
)

type Controller struct {
	config util.Config
	store  db.Querier
}

func NewController(config util.Config, store db.Querier) *Controller {
	return &Controller{
		config: config,
		store:  store,
	}
}

func (c *Controller) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Get("/csv", c.getCsv)
	router.Get("/mean", c.getMean)
	router.Get("/data", c.getData)

	return router
}
