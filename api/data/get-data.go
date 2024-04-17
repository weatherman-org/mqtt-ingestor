package data

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/weatherman-org/telemetry/db/sqlc"
	"github.com/weatherman-org/telemetry/util"
)

type getDataResponse struct {
	Millis        int64   `json:"millis"`
	Temperature   float64 `json:"temperature"`
	Humidity      float64 `json:"humidity"`
	WindSpeed     float64 `json:"wind_speed"`
	WindDirection float64 `json:"wind_direction"`
	Pressure      float64 `json:"pressure"`
	WaterAmount   float64 `json:"water_amount"`
}

// getData godoc
// @Summary      Get individual data for a ranged time frame
// @Description  Get individual data for a ranged time frame
// @Tags         data
// @Produce      json
// @Param        start query string true "json"
// @Param        end query string true "json"
// @Success      200  {object} []getDataResponse
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /data/data [get]
func (c *Controller) getData(w http.ResponseWriter, r *http.Request) {
	startStr, endStr := r.URL.Query().Get("start"), r.URL.Query().Get("end")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	end, err := strconv.Atoi(endStr)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	arg := db.GetWeatherTelemetryByRangeParams{
		Millis:   time.UnixMilli(int64(start)),
		Millis_2: time.UnixMilli(int64(end)),
	}
	data, err := c.store.GetWeatherTelemetryByRange(r.Context(), arg)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	res := make([]getDataResponse, len(data))
	for d := range data {
		res[d] = getDataResponse{
			Millis:        data[d].Millis.UnixMilli(),
			Temperature:   data[d].Temperature,
			Humidity:      data[d].Humidity,
			WindSpeed:     data[d].Windspeed,
			WindDirection: data[d].Winddirection,
			Pressure:      data[d].Pressure,
			WaterAmount:   data[d].Wateramount,
		}
	}

	util.WriteJson(w, http.StatusOK, res)
}
