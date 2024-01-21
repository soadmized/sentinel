package service

import (
	"time"

	"github.com/pkg/errors"

	"sentinel/internal/dataset"
)

type Repository interface {
	SaveValues(dataset dataset.Dataset) error

	LastValues(sensorID string) *dataset.Dataset
	Values(sensorID string, start, end time.Duration) *dataset.Dataset
}

type Service struct {
	Repo Repository
}

func (s *Service) SaveValues(dataset dataset.Dataset) error {
	err := s.Repo.SaveValues(dataset)
	if err != nil {
		return errors.Wrap(err, "save values")
	}

	return nil
}

func (s *Service) LastValues(sensorID string) *dataset.Dataset {
	return s.Repo.LastValues(sensorID)
}
