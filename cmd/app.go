package cmd

import (
	"log"
	"sentinel/internal/build"
	"sentinel/internal/config"
)

func Run(conf config.Config) error {
	builder, err := build.New(conf)
	if err != nil {
		return err
	}

	a, err := builder.Api()
	if err != nil {
		return err
	}

	a.Route()

	log.Fatal(a.Start(conf.AppPort))

	return nil
}
