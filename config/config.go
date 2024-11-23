package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	URL_PORT string

	DATABASE_HOST     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	DATABASE_PORT     string
	DATABASE_SSL      string
	TOKEN_LOGIN       string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}
