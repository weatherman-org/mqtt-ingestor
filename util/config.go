package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MQTT_ENDPOINT string `mapstructure:"MQTT_ENDPOINT"`
	MQTT_USERNAME string `mapstructure:"MQTT_USERNAME"`
	MQTT_PASSWORD string `mapstructure:"MQTT_PASSWORD"`
	HTTP_PORT     string `mapstructure:"HTTP_PORT"`
	POSTGRES_DSN  string `mapstructure:"POSTGRES_DSN"`
	MIGRATION_URL string `mapstructure:"MIGRATION_URL"`
}

func LoadEnv(path string) (config Config, err error) {
	err = godotenv.Load(path + "/.env")

	config = Config{
		MQTT_ENDPOINT: os.Getenv("MQTT_ENDPOINT"),
		MQTT_USERNAME: os.Getenv("MQTT_USERNAME"),
		MQTT_PASSWORD: os.Getenv("MQTT_PASSWORD"),
		HTTP_PORT:     os.Getenv("HTTP_PORT"),
		POSTGRES_DSN:  os.Getenv("POSTGRES_DSN"),
		MIGRATION_URL: os.Getenv("MIGRATION_URL"),
	}
	return
}
