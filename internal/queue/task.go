package queue

import (
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"

	"github.com/soadmized/sentinel/pkg/dataset"
)

type SingleDatasetTask struct {
	Dataset dataset.Dataset `json:"dataset"`
}

type MultipleDatasetTask struct {
	Sets []dataset.Dataset `json:"sets"`
}

func NewWriteValuesTask(values dataset.Dataset) (*asynq.Task, error) {
	payload, err := json.Marshal(SingleDatasetTask{Dataset: values})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeSingleDatasetTask, payload), nil
}

// AggregateTasks aggregates multiple single task into 1
func AggregateTasks(group string, tasks []*asynq.Task) *asynq.Task {
	datasets := make([]dataset.Dataset, 0, len(tasks))

	log.Print("start aggregating single tasks")

	for _, t := range tasks {
		var payload SingleDatasetTask

		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			log.Printf("unmarshal task payload: %v", err)

			continue
		}

		datasets = append(datasets, payload.Dataset)
	}

	buf, err := json.Marshal(MultipleDatasetTask{Sets: datasets})
	if err != nil {
		log.Printf("can't marshal task: %v", err)
	}

	log.Print("aggregating was completed")

	return asynq.NewTask(TypeMultipleDatasetTask, buf)
}
