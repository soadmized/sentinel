package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort int    `envconfig:"APP_PORT" default:"8080"`
	AppUser string `envconfig:"APP_USER"`
	AppPass string `envconfig:"APP_PASS"`

	RedisHost string `envconfig:"REDIS_HOST"`

	Influx Influx `envconfig:"INFLUX"`
	Kafka  Kafka  `envconfig:"KAFKA"`

	PostgresDSN string `envconfig:"POSTGRES_DSN"`
}

type Influx struct {
	Port   int    `envconfig:"PORT" default:"8086"`
	Token  string `envconfig:"TOKEN"`
	Org    string `envconfig:"ORG" default:"meltdown"`
	Bucket string `envconfig:"BUCKET" default:"super-bucket"`
}

type Kafka struct {
	Brokers      []string `envconfig:"BROKERS"`
	User         string   `envconfig:"USER"`
	Pass         string   `envconfig:"PASS"`
	UserGroup    string   `envconfig:"USER_GROUP"`
	DatasetTopic string   `envconfig:"TOPIC_DATASET"` // produce
	EventTopic   string   `envconfig:"TOPIC_EVENT"`   // consume
}

func Read() Config {
	conf := Config{} //nolint:exhaustruct

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}

	return conf
}
