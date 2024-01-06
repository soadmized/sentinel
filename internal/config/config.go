package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort int    `envconfig:"APP_PORT" default:"8080"`
	Influx  Influx `envconfig:"INFLUX"`
}

type Influx struct {
	Port   string `envconfig:"PORT" default:"8086"`
	Token  string `envconfig:"TOKEN"`
	Org    string `envconfig:"ORG" default:"meltdown"`
	Bucket string `envconfig:"BUCKET" default:"super-bucket"`
}

func Read() Config {
	conf := Config{}

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}

	return conf
}
