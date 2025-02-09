package producer

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/soadmized/sentinel/pkg/dataset"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Client struct {
	client *kgo.Client
	topic  string
}

func New(cl *kgo.Client, topic string) *Client {
	return &Client{
		client: cl,
		topic:  topic,
	}
}

func (c *Client) ProduceDataset(ctx context.Context, payload dataset.Dataset) error {
	rec, err := c.datasetToRecord(payload)
	if err != nil {
		return errors.Wrap(err, "convert dataset to record")
	}

	wg := sync.WaitGroup{}

	wg.Add(1)

	c.client.Produce(ctx, rec, func(r *kgo.Record, err error) {
		defer wg.Done()

		if err != nil {
			log.Printf("record had a produce error: %v\n", err)
		}
	})

	wg.Wait()

	return nil
}

func (c *Client) ConsumeEvent(ctx context.Context) error {
	return nil
}

func (c *Client) datasetToRecord(ds dataset.Dataset) (*kgo.Record, error) {
	b, err := json.Marshal(record(ds))
	if err != nil {
		return nil, errors.Wrap(err, "marshal dataset")
	}

	rec := &kgo.Record{
		Key:   []byte(uuid.New().String()),
		Value: b,
		Topic: c.topic,
	}

	return rec, nil
}
