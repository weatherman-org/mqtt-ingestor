package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/weathermamn-org/telemetry/data"
	db "github.com/weathermamn-org/telemetry/db/sqlc"
	"github.com/weathermamn-org/telemetry/mqtt"
	"github.com/weathermamn-org/telemetry/util"
	"google.golang.org/protobuf/proto"
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

func (s *Server) addRoutes() {
	router := chi.NewMux()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/ping"),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	router.Post("/publish", s.handlePublish)

	s.router = router
}

type publishInput struct {
	Topic         string  `json:"topic" validate:"required"`
	Temperature   float64 `json:"temperature"`
	Humidity      float64 `json:"humidity"`
	WindSpeed     float64 `json:"wind_speed"`
	WindDirection float64 `json:"wind_direction"`
	Pressure      float64 `json:"pressure"`
	WaterAmount   float64 `json:"water_amount"`
}

func (s *Server) handlePublish(w http.ResponseWriter, r *http.Request) {
	var requestPayload publishInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	// * filliing default values if zero
	fillDefaultIfZero(&requestPayload.Temperature)
	fillDefaultIfZero(&requestPayload.Humidity)
	fillDefaultIfZero(&requestPayload.WindSpeed)
	fillDefaultIfZero(&requestPayload.WindDirection)
	fillDefaultIfZero(&requestPayload.Pressure)
	fillDefaultIfZero(&requestPayload.WaterAmount)

	weatherProto := &data.WeatherTelemetry{
		Timestamp:     uint64(time.Now().UnixMilli()),
		Temperature:   requestPayload.Temperature,
		Humidity:      requestPayload.Humidity,
		WindSpeed:     requestPayload.WindSpeed,
		WindDirection: requestPayload.WindDirection,
		Pressure:      requestPayload.Pressure,
		WaterAmount:   requestPayload.WaterAmount,
	}

	p, err := proto.Marshal(weatherProto)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	if err := s.mqttSession.Publish(p, requestPayload.Topic); err != nil {
		util.ErrorJson(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, requestPayload)
}

func fillDefaultIfZero(field *float64) {
	if *field == 0 {
		*field = util.RandomFloat64(0, 100)
	}
}
