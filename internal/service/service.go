package service

import (
	"sentry/internal/repo"
)

type Repo interface {
	SaveTemperature(temp float32) error
	SaveMotion(motion bool) error
	SaveLight(light int) error

	GetLastValues() *repo.DataSet
}

type Service struct {
	Repo Repo
}
