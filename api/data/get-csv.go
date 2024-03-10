package data

import (
	"bytes"
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"github.com/weatherman-org/telemetry/util"
)

// getCsv godoc
// @Summary      Get a CSV formatted file of all the MQTT data
// @Description  Get a CSV formatted file of all the MQTT data
// @Tags         data
// @Produce      text/csv
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /data/csv [get]
func (c *Controller) getCsv(w http.ResponseWriter, r *http.Request) {
	var csvBuffer bytes.Buffer
	writer := csv.NewWriter(&csvBuffer)
	defer writer.Flush()

	header := []string{"timestamp", "temperature", "humidity", "wind speed", "wind direction", "pressure", "water amount"}
	if err := writer.Write(header); err != nil {
		util.ErrorJson(w, err)
		return
	}

	dataCompleted := false
	var startTimestamp int64 = 0
	for !dataCompleted {
		telemetry, err := c.store.GetWeatherTelemetry(r.Context(), time.UnixMilli(startTimestamp))
		if err != nil {
			util.ErrorJson(w, err)
			return
		}

		if len(telemetry) == 0 {
			dataCompleted = true
			break
		}

		for _, t := range telemetry {
			record := []string{
				strconv.FormatInt(t.Millis.UnixMilli(), 10),
				strconv.FormatFloat(t.Temperature, 'f', 2, 64),
				strconv.FormatFloat(t.Humidity, 'f', 2, 64),
				strconv.FormatFloat(t.Windspeed, 'f', 2, 64),
				strconv.FormatFloat(t.Winddirection, 'f', 2, 64),
				strconv.FormatFloat(t.Pressure, 'f', 2, 64),
				strconv.FormatFloat(t.Wateramount, 'f', 2, 64),
			}
			writer.Write(record)
		}

		startTimestamp = telemetry[len(telemetry)-1].Millis.UnixMilli()
	}

	writer.Flush()

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=weather-data.csv")

	w.Write(csvBuffer.Bytes())
}
