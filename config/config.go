package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env            string
	GIN_MODE       string `envconfig:"GIN_MODE" default:"release"`
	MYSQL_HOSTNAME string `envconfig:"MYSQL_HOSTNAME"`
	MYSQL_PORT     string `envconfig:"MYSQL_PORT"`
	MYSQL_USERNAME string `envconfig:"MYSQL_USERNAME"`
	MYSQL_PASSWORD string `envconfig:"MYSQL_PASSWORD"`
	MYSQL_DATABASE string `envconfig:"MYSQL_DATABASE"`
}

func Load() Config {
	var config Config

	ENV, ok := os.LookupEnv("ENV")

	if !ok {
		// Default value for ENV.
		ENV = "dev"
	}

	err := godotenv.Load("./.env")
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	envconfig.MustProcess("", &config)
	config.Env = ENV
	return config
}
