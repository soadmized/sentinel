package service

import (
	"context"
	"fmt"
	"time"

	"github.com/soadmized/sentinel/pkg/dataset"
)

type Repository interface {
	SaveLastValues(dataset dataset.Dataset) error
	LastValues(sensorID string) *dataset.Dataset
	SensorIDs() []string
}

type QueueClient interface {
	Enqueue(values dataset.Dataset) error
}

type Service struct {
	Repo        Repository
	QueueClient QueueClient
}

// SaveValues stores values from sensors in repo.
func (s *Service) SaveValues(_ context.Context, dataset dataset.Dataset) error {
	if err := s.Repo.SaveLastValues(dataset); err != nil {
		return fmt.Errorf("save values: %w", err)
	}

	if err := s.QueueClient.Enqueue(dataset); err != nil {
		return fmt.Errorf("add task to queue: %w", err)
	}

	return nil
}

func (s *Service) SensorIDs() []string {
	return s.Repo.SensorIDs()
}

// LastValues returns last values from repo for given sensorID.
func (s *Service) LastValues(_ context.Context, sensorID string) *dataset.Dataset {
	return s.Repo.LastValues(sensorID)
}

// SensorStatuses returns sensorIDs and true if sensor was online in last 10 seconds, else false.
func (s *Service) SensorStatuses(ctx context.Context) map[string]bool {
	ids := s.Repo.SensorIDs()
	statuses := make(map[string]bool)

	for _, id := range ids {
		lastValues := s.LastValues(ctx, id)
		lastSeen := lastValues.UpdatedAt
		online := checkIfOnline(lastSeen)
		statuses[id] = online
	}

	return statuses
}

func checkIfOnline(lastSeen time.Time) bool {
	return time.Since(lastSeen) <= time.Second*10
}
