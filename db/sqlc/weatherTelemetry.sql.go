// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: weatherTelemetry.sql

package db

import (
	"context"
	"time"
)

const insertWeatherTelemetry = `-- name: InsertWeatherTelemetry :one
INSERT INTO weatherTelemetry(
        millis,
        temperature,
        humidity,
        windSpeed,
        windDirection,
        pressure,
        waterAmount
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING millis, temperature, humidity, windspeed, winddirection, pressure, wateramount
`

type InsertWeatherTelemetryParams struct {
	Millis        time.Time `json:"millis"`
	Temperature   float64   `json:"temperature"`
	Humidity      float64   `json:"humidity"`
	Windspeed     float64   `json:"windspeed"`
	Winddirection float64   `json:"winddirection"`
	Pressure      float64   `json:"pressure"`
	Wateramount   float64   `json:"wateramount"`
}

func (q *Queries) InsertWeatherTelemetry(ctx context.Context, arg InsertWeatherTelemetryParams) (Weathertelemetry, error) {
	row := q.db.QueryRow(ctx, insertWeatherTelemetry,
		arg.Millis,
		arg.Temperature,
		arg.Humidity,
		arg.Windspeed,
		arg.Winddirection,
		arg.Pressure,
		arg.Wateramount,
	)
	var i Weathertelemetry
	err := row.Scan(
		&i.Millis,
		&i.Temperature,
		&i.Humidity,
		&i.Windspeed,
		&i.Winddirection,
		&i.Pressure,
		&i.Wateramount,
	)
	return i, err
}