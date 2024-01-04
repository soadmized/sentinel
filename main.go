package main

import (
	"log"

	"sentry/cmd"
	"sentry/internal/config"
)

func main() {
	conf := config.Read()

	err := cmd.Run(conf)
	if err != nil {
		log.Fatal(err)
	}
}
