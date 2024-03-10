package cmd

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/weathermamn-org/mqtt-ingestor/mqtt"
	"github.com/weathermamn-org/mqtt-ingestor/util"
)

func Subscribe(config util.Config) error {
	session, err := setup(config)
	if err != nil {
		return err
	}

	return session.Subscribe("topic/telemetry", func(client paho.Client, message paho.Message) {
		log.Println(string(message.Payload()))
	})
}

func setup(config util.Config) (mqtt.Session, error) {
	return mqtt.NewSession(config.MQTT_ENDPOINT, config.MQTT_USERNAME, config.MQTT_PASSWORD)
}
