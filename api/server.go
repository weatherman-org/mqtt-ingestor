package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/weatherman-org/telemetry/api/data"
	db "github.com/weatherman-org/telemetry/db/sqlc"
	"github.com/weatherman-org/telemetry/mqtt"
	"github.com/weatherman-org/telemetry/util"
)

type Server struct {
	config      util.Config
	mqttSession mqtt.Session
	store       db.Querier

	dataController *data.Controller
	router         *chi.Mux
}

func NewServer(config util.Config, session mqtt.Session, store db.Querier) *Server {
	s := &Server{
		config:      config,
		mqttSession: session,
		store:       store,
	}

	s.dataController = data.NewController(s.config, s.store)

	s.addRoutes()

	return s
}

func (s *Server) Start() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.HTTP_PORT),
		Handler: s.router,
	}
	return srv.ListenAndServe()
}
