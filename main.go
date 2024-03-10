package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/weatherman-org/telemetry/api"
	db "github.com/weatherman-org/telemetry/db/sqlc"
	"github.com/weatherman-org/telemetry/mqtt"
	"github.com/weatherman-org/telemetry/util"
)

// @title Weatherman Telemetry Server APIs
// @version 1.0
// @description The comprehensive list of all Weatherman Telemetry Server APIs
// @host localhost:8080
// @BasePath /
func main() {
	config, err := util.LoadEnv(".")
	if err != nil {
		log.Println("unable to load env:", err)
	}

	conn := util.CreatePostgresPool(config.POSTGRES_DSN)
	defer conn.Close()

	util.CreateDatabase(conn)

	fmt.Print("attempting database migration...")
	if err := runDbMigration(config); err != nil {
		fmt.Println(" database migration failed with error:", err)
	} else {
		fmt.Println(" database migration was successful!")
	}

	store := db.New(conn)

	session, err := mqtt.NewSession(config.MQTT_ENDPOINT, config.MQTT_USERNAME, config.MQTT_PASSWORD)
	if err != nil {
		log.Panic("unable to create mqtt session:", err)
	}
	defer session.Disconnect(1000)

	httpServer := api.NewServer(config, session, store)
	log.Println("starting the http server on port", config.HTTP_PORT, "...")
	go httpServer.Start()

	log.Println("starting the mqtt subscriber...")
	if err := mqtt.Subscribe(session, store); err != nil {
		log.Panic("unable to subscribe: ", err)
	}
}

func runDbMigration(config util.Config) error {
	migration, err := migrate.New(config.MIGRATION_URL, config.POSTGRES_DSN)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
