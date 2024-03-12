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
LIMIT $2;
-- name: GetMeanWeatherTelemetry :one
SELECT AVG(temperature) AS temperature,
    AVG(humidity) AS humidity,
    AVG(windSpeed) AS windSpeed,
    AVG(windDirection) AS windDirection,
    AVG(pressure) AS pressure,
    AVG(waterAmount) AS waterAmount
FROM weatherTelemetry
WHERE millis >= $1
    and millis < $2
LIMIT 1;