package cmd

import (
	"log"

	"github.com/pkg/errors"

	"github.com/soadmized/sentinel/internal/build"
	"github.com/soadmized/sentinel/internal/config"
)

func Run(conf config.Config) error {
	builder, err := build.New(conf)
	if err != nil {
		return errors.Wrap(err, "build is failed")
	}

	a, err := builder.API()
	if err != nil {
		return errors.Wrap(err, "get api is failed")
	}

	a.Route()

	log.Fatal(a.Start(conf.AppPort))

	return nil
}
