package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"

	"github.com/soadmized/sentinel/pkg/dataset"
)

const (
	TypeSingleDatasetTask   = "single-dataset"
	TypeMultipleDatasetTask = "multiple-dataset"

	GroupSingleDatasets = "single-datasets"

	MainQueue  = "main"
	RetryQueue = "retry"
)

type repo interface {
	SaveValues(ctx context.Context, dataset dataset.Dataset) error
}

type Queue struct {
	repo repo
}

func New(repo repo) Queue {
	return Queue{repo: repo}
}

// ProcessTask handles task with multiple datasets
func (q *Queue) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var writeTask MultipleDatasetTask

	if err := json.Unmarshal(task.Payload(), &writeTask); err != nil {
		return fmt.Errorf("unmarshal task payload: %w", err)
	}

	for _, set := range writeTask.Sets {
		if err := q.repo.SaveValues(ctx, set); err != nil {
			return fmt.Errorf("save values to repo: %w", err)
		}
	}

	return nil
}
