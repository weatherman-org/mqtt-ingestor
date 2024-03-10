-- name: InsertWeatherTelemetry :one
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
RETURNING *;
-- name: GetWeatherTelemetry :many
SELECT *
FROM weatherTelemetry
WHERE millis > $1
ORDER BY millis ASC
LIMIT 100;