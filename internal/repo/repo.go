package repo

import (
	"context"
	"log"
	"sync"
	"time"

	influxclient "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"

	"github.com/soadmized/sentinel/internal/repo/faststorage"
	"github.com/soadmized/sentinel/pkg/dataset"
)

type Repo struct {
	fastStorage fstorage.FastStorage
	writer      influxapi.WriteAPI
	reader      influxapi.QueryAPI
}

func New(writer influxapi.WriteAPI, reader influxapi.QueryAPI) (Repo, error) {
	storage := make(fstorage.FastStorage)

	return Repo{
		fastStorage: storage,
		writer:      writer,
		reader:      reader,
	}, nil
}

func (r *Repo) SaveLastValues(dataset dataset.Dataset) error {
	r.fastStorage[dataset.ID] = dataset

	return nil
}

func (r *Repo) SaveValues(_ context.Context, dataset dataset.Dataset) error {
	wg := sync.WaitGroup{}
	wg.Add(3) //nolint:gomnd

	go func() {
		r.saveTemperature(dataset.ID, dataset.UpdatedAt, dataset.Temp)

		defer wg.Done()
	}()

	go func() {
		r.saveMotion(dataset.ID, dataset.UpdatedAt, dataset.Motion)

		defer wg.Done()
	}()

	go func() {
		r.saveLight(dataset.ID, dataset.UpdatedAt, dataset.Light)

		defer wg.Done()
	}()

	log.Print("dataset was save to influx")

	return nil
}

func (r *Repo) saveTemperature(id string, stamp time.Time, temp float32) {
	point := influxclient.NewPointWithMeasurement("temperature").
		AddTag("id", id).
		AddField("temperature", temp).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

func (r *Repo) saveMotion(id string, stamp time.Time, motion bool) {
	point := influxclient.NewPointWithMeasurement("motion").
		AddTag("id", id).
		AddField("motion", motion).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

func (r *Repo) saveLight(id string, stamp time.Time, light int) {
	point := influxclient.NewPointWithMeasurement("light").
		AddTag("id", id).
		AddField("light", light).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

// LastValues gets last stored values.
func (r *Repo) LastValues(sensorID string) *dataset.Dataset {
	values, ok := r.fastStorage[sensorID]
	if !ok {
		return nil
	}

	return &values
}

// Values gets set of values between start and end.
func (r *Repo) Values(sensorID string, start, end time.Time) *dataset.Dataset {
	// TODO temporary implementation
	values, ok := r.fastStorage[sensorID]
	if !ok {
		return nil
	}

	return &values
}

// SensorIDs returns all sensorIDs stored in fast storage.
func (r *Repo) SensorIDs() []string {
	ids := make([]string, 0, len(r.fastStorage))

	for id := range r.fastStorage {
		ids = append(ids, id)
	}

	return ids
}
