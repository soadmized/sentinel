package service

import (
	"sentinel/internal/model"
	"time"

	"github.com/pkg/errors"
)

type Repository interface {
	SaveValues(dataset model.Dataset) error

	LastValues() *model.Dataset
	Values(start, end time.Duration) *model.Dataset
}

type Service struct {
	Repo Repository
}

func (s *Service) SaveValues(dataset model.Dataset) error {
	err := s.Repo.SaveValues(dataset)
	if err != nil {
		return errors.Wrap(err, "save values")
	}

	return nil
}

func (s *Service) LastValues() *model.Dataset {
	return s.Repo.LastValues()
}
