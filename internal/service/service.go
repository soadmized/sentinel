package service

import (
	"sentry/internal/repo"
	"time"
)

type Repository interface {
	SaveValues(temp float32, motion bool, light int) error

	LastValues() *repo.DataSet
	Values(start, end time.Duration) *repo.DataSet
}

type Service struct {
	Repo Repository
}
