package service

import (
	"time"

	"github.com/pkg/errors"
	"github.com/soadmized/sentinel/pkg/dataset"
)

type Repository interface {
	SaveValues(dataset dataset.Dataset) error
	LastValues(sensorID string) *dataset.Dataset
	Values(sensorID string, start, end time.Duration) *dataset.Dataset
	SensorIDs() []string
}

type Service struct {
	Repo Repository
}

// SaveValues stores values from sensors in repo.
func (s *Service) SaveValues(dataset dataset.Dataset) error {
	err := s.Repo.SaveValues(dataset)
	if err != nil {
		return errors.Wrap(err, "save values")
	}

	return nil
}

func (s *Service) SensorIDs() []string {
	return s.Repo.SensorIDs()
}

// LastValues returns last values from repo for given sensorID.
func (s *Service) LastValues(sensorID string) *dataset.Dataset {
	return s.Repo.LastValues(sensorID)
}

// SensorStatuses returns sensorIDs and true if sensor was online in last 10 seconds, else false.
func (s *Service) SensorStatuses() map[string]bool {
	ids := s.Repo.SensorIDs()
	statuses := make(map[string]bool)

	for _, id := range ids {
		lastValues := s.LastValues(id)
		lastSeen := lastValues.UpdatedAt
		online := checkIfOnline(lastSeen)
		statuses[id] = online
	}

	return statuses
}

func checkIfOnline(lastSeen time.Time) bool {
	return time.Since(lastSeen) <= time.Second*10
}
