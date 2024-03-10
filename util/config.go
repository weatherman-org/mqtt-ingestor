package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	MQTT_ENDPOINT string `mapstructure:"MQTT_ENDPOINT"`
	MQTT_USERNAME string `mapstructure:"MQTT_USERNAME"`
	MQTT_PASSWORD string `mapstructure:"MQTT_PASSWORD"`
}

func LoadEnv(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}
