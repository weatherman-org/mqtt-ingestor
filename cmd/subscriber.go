package cmd

import (
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/weathermamn-org/mqtt-ingestor/data"
	"github.com/weathermamn-org/mqtt-ingestor/mqtt"
	"google.golang.org/protobuf/proto"
)

func Subscribe(session mqtt.Session) error {
	return session.Subscribe("topic/telemetry", func(client paho.Client, message paho.Message) {
		messageValue := message.Payload()

		// * attempting unmarshal of message
		var weatherData data.WeatherTelemetry
		if err := proto.Unmarshal(messageValue, &weatherData); err != nil {
			log.Println("unable to unmarshal message:", err)
			log.Println("message was:", string(messageValue))
		} else {
			j, _ := json.MarshalIndent(weatherData, "", "\t")
			log.Println(string(j))
		}
	})
}
