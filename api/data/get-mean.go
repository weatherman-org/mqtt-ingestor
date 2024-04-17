package data

import (
	"context"
	"net/http"
	"time"

	"github.com/papaya147/parallelize"
	db "github.com/weatherman-org/telemetry/db/sqlc"
	"github.com/weatherman-org/telemetry/util"
)

type mean struct {
	MeanTemperature   float64 `json:"mean_temperature"`
	MeanHumidity      float64 `json:"mean_humidity"`
	MeanWindSpeed     float64 `json:"mean_wind_speed"`
	MeanWindDirection float64 `json:"mean_wind_direction"`
	MeanPressure      float64 `json:"mean_pressure"`
	MeanWaterAmount   float64 `json:"mean_water_amount"`
}

type getMeanResponse struct {
	DailyMean   mean `json:"daily_mean"`
	WeeklyMean  mean `json:"weekly_mean"`
	MonthlyMean mean `json:"monthly_mean"`
	YearlyMean  mean `json:"yearly_mean"`
}

// getMean godoc
// @Summary      Get mean data for different time frames
// @Description  Get mean data for daily, weekly, monthly and yearly time frames
// @Tags         data
// @Produce      json
// @Success      200  {object} getMeanResponse
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /data/mean [get]
func (c *Controller) getMean(w http.ResponseWriter, r *http.Request) {
	group := parallelize.NewSyncGroup()
	var dailyMean mean
	parallelize.AddOutputtingMethodWithArgs(group, c.selectMeanData, parallelize.OutputtingMethodWithArgsParams[string, *mean]{
		Context: r.Context(),
		Input:   "day",
		Output:  &dailyMean,
	})

	var weeklyMean mean
	parallelize.AddOutputtingMethodWithArgs(group, c.selectMeanData, parallelize.OutputtingMethodWithArgsParams[string, *mean]{
		Context: r.Context(),
		Input:   "week",
		Output:  &weeklyMean,
	})

	var monthlyMean mean
	parallelize.AddOutputtingMethodWithArgs(group, c.selectMeanData, parallelize.OutputtingMethodWithArgsParams[string, *mean]{
		Context: r.Context(),
		Input:   "month",
		Output:  &monthlyMean,
	})

	var yearlyMean mean
	parallelize.AddOutputtingMethodWithArgs(group, c.selectMeanData, parallelize.OutputtingMethodWithArgsParams[string, *mean]{
		Context: r.Context(),
		Input:   "year",
		Output:  &yearlyMean,
	})

	if err := group.Run(); err != nil {
		util.ErrorJson(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, getMeanResponse{
		DailyMean:   dailyMean,
		WeeklyMean:  weeklyMean,
		MonthlyMean: monthlyMean,
		YearlyMean:  yearlyMean,
	})
}

func (c *Controller) selectMeanData(ctx context.Context, timeframe string, out *mean) error {
	var start, end time.Time
	switch timeframe {
	case "day":
		start = time.Now().AddDate(0, 0, -1)
		end = time.Now()
	case "week":
		start = time.Now().AddDate(0, 0, -7)
		end = time.Now()
	case "month":
		start = time.Now().AddDate(0, -1, 0)
		end = time.Now()
	default:
		start = time.Now().AddDate(-1, 0, 0)
		end = time.Now()
	}

	args := db.GetMeanWeatherTelemetryParams{
		Millis:   start,
		Millis_2: end,
	}
	meanTelemetry, err := c.store.GetMeanWeatherTelemetry(ctx, args)
	if err != nil {
		return err
	}

	out.MeanTemperature = meanTelemetry.Temperature
	out.MeanHumidity = meanTelemetry.Humidity
	out.MeanWindSpeed = meanTelemetry.Windspeed
	out.MeanWindDirection = meanTelemetry.Winddirection
	out.MeanPressure = meanTelemetry.Pressure
	out.MeanWaterAmount = meanTelemetry.Wateramount

	return nil
}
