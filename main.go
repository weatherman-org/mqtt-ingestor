package main

import (
	"log"

	"github.com/weathermamn-org/mqtt-ingestor/cmd"
	"github.com/weathermamn-org/mqtt-ingestor/mqtt"
	"github.com/weathermamn-org/mqtt-ingestor/util"
)

func main() {
	config, err := util.LoadEnv(".")
	if err != nil {
		log.Println("unable to load env:", err)
	}

	session, err := mqtt.NewSession(config.MQTT_ENDPOINT, config.MQTT_USERNAME, config.MQTT_PASSWORD)
	if err != nil {
		log.Panic("unable to create mqtt session:", err)
	}
	defer session.Disconnect(1000)

	httpServer := cmd.NewServer(config, session)
	log.Println("started the http server on port", config.HTTP_PORT)
	go httpServer.Start()

	if err := cmd.Subscribe(session); err != nil {
		log.Panic("unable to subscribe: ", err)
	} else {
		log.Println("started the mqtt subscriber")
	}
}
