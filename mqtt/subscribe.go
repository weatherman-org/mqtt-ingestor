package mqtt

import (
	"context"
	"encoding/json"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	db "github.com/weatherman-org/telemetry/db/sqlc"
	"github.com/weatherman-org/telemetry/weatherdata"
	"google.golang.org/protobuf/proto"
)

func Subscribe(session Session, store db.Querier) error {
	return session.Subscribe("topic/telemetry", func(client paho.Client, message paho.Message) {
		messageValue := message.Payload()

		// * attempting unmarshal of message
		var weatherData weatherdata.WeatherTelemetry
		if err := proto.Unmarshal(messageValue, &weatherData); err != nil {
			log.Println("unable to unmarshal message:", err)
			log.Println("message was:", string(messageValue))
		} else {
			j, _ := json.MarshalIndent(weatherData, "", "\t")
			log.Println(string(j))
			if _, err := store.InsertWeatherTelemetry(context.Background(), db.InsertWeatherTelemetryParams{
				Millis:        time.UnixMilli(int64(weatherData.Timestamp)),
				Temperature:   weatherData.Temperature,
				Humidity:      weatherData.Humidity,
				Windspeed:     weatherData.WindSpeed,
				Winddirection: weatherData.WindDirection,
				Pressure:      weatherData.Pressure,
				Wateramount:   weatherData.WaterAmount,
			}); err != nil {
				log.Println("error inserting data to store:", err)
			}
		}
	})
}
