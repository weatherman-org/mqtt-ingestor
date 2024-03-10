package api

import (
	"net/http"
	"time"

	"github.com/weatherman-org/telemetry/data"
	"github.com/weatherman-org/telemetry/util"
	"google.golang.org/protobuf/proto"
)

type publishModel struct {
	Topic         string  `json:"topic" validate:"required"`
	Temperature   float64 `json:"temperature"`
	Humidity      float64 `json:"humidity"`
	WindSpeed     float64 `json:"wind_speed"`
	WindDirection float64 `json:"wind_direction"`
	Pressure      float64 `json:"pressure"`
	WaterAmount   float64 `json:"water_amount"`
}

// create godoc
// @Summary      Mock an MQTT weather publish
// @Description  Mock an MQTT weather publish, fields are formatted as Protobuf and sent via the MQTT broker
// @Tags         mqtt
// @Param        request-body body publishModel true "json"
// @Success      200  {object} publishModel
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /publish [post]
func (s *Server) mqttPublish(w http.ResponseWriter, r *http.Request) {
	var requestPayload publishModel
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
