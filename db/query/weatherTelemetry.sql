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
SELECT CAST(
        COALESCE(AVG(temperature), 0) AS DOUBLE PRECISION
    ) AS temperature,
    CAST(COALESCE(AVG(humidity), 0) AS DOUBLE PRECISION) AS humidity,
    CAST(COALESCE(AVG(windSpeed), 0) AS DOUBLE PRECISION) AS windSpeed,
    CAST(
        COALESCE(AVG(windDirection), 0) AS DOUBLE PRECISION
    ) AS windDirection,
    CAST(COALESCE(AVG(pressure), 0) AS DOUBLE PRECISION) AS pressure,
    CAST(
        COALESCE(AVG(waterAmount), 0) AS DOUBLE PRECISION
    ) AS waterAmount
FROM weatherTelemetry
WHERE millis >= $1
    and millis < $2
LIMIT 1;
-- name: GetWeatherTelemetryByRange :many
SELECT *
FROM weatherTelemetry
WHERE millis >= $1
    AND millis < $2
ORDER BY millis ASC;