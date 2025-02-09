package qclient

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"

	"github.com/soadmized/sentinel/internal/queue"
	"github.com/soadmized/sentinel/pkg/dataset"
)

type Client struct {
	client *asynq.Client
}

func New(asynqClient *asynq.Client) *Client {
	return &Client{client: asynqClient}
}

func (c *Client) Enqueue(values dataset.Dataset) error {
	task, err := queue.NewWriteValuesTask(values)
	if err != nil {
		return fmt.Errorf("create new task with values %v: %w", values, err)
	}

	info, err := c.client.Enqueue(task, asynq.Queue(queue.MainQueue), asynq.Group(queue.GroupSingleDatasets))
	if err != nil {
		return fmt.Errorf("enqueue task with values %v: %w", values, err)
	}

	log.Printf("single task %s was added to %s queue and %s group", info.ID, info.Queue, info.Group)

	return nil
}
