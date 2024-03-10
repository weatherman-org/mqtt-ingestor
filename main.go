package main

import (
	"log"

	"github.com/weathermamn-org/mqtt-ingestor/cmd"
	"github.com/weathermamn-org/mqtt-ingestor/util"
)

func main() {
	config, err := util.LoadEnv(".")
	if err != nil {
		log.Panic("unable to load env: ", err)
	}

	if err := cmd.Subscribe(config); err != nil {
		log.Panic("unable to subscribe: ", err)
	}
}
