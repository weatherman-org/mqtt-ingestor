package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	db "github.com/weathermamn-org/telemetry/db/sqlc"
	"github.com/weathermamn-org/telemetry/mqtt"
	"github.com/weathermamn-org/telemetry/util"
)

type Server struct {
	config      util.Config
	mqttSession mqtt.Session
	store       db.Querier
	router      *chi.Mux
}

func NewServer(config util.Config, session mqtt.Session, store db.Querier) *Server {
	s := &Server{
		config:      config,
		mqttSession: session,
		store:       store,
	}

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
