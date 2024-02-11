package repo

import (
	"fmt"
	"sync"
	"time"

	influxClient "github.com/influxdata/influxdb-client-go/v2"
	influxApi "github.com/influxdata/influxdb-client-go/v2/api"

	"sentinel/internal/config"
	"sentinel/internal/dataset"
	"sentinel/internal/repo/fstorage"
)

type Repo struct {
	fastStorage fstorage.FastStorage
	writer      influxApi.WriteAPI
	reader      influxApi.QueryAPI
}

func New(conf config.Config) (Repo, error) {
	storage := make(fstorage.FastStorage)

	url := fmt.Sprintf("http://localhost:%d", conf.Influx.Port)
	client := influxClient.NewClient(url, conf.Influx.Token)

	writer := client.WriteAPI(conf.Influx.Org, conf.Influx.Bucket)
	reader := client.QueryAPI(conf.Influx.Org)

	return Repo{
		fastStorage: storage,
		writer:      writer,
		reader:      reader,
	}, nil
}

func (r *Repo) SaveValues(dataset dataset.Dataset) error {
	r.fastStorage[dataset.ID] = dataset

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

	return nil
}

func (r *Repo) saveTemperature(id string, stamp time.Time, temp float32) {
	point := influxClient.NewPointWithMeasurement("temperature").
		AddTag("id", id).
		AddField("temperature", temp).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

func (r *Repo) saveMotion(id string, stamp time.Time, motion bool) {
	point := influxClient.NewPointWithMeasurement("motion").
		AddTag("id", id).
		AddField("motion", motion).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

func (r *Repo) saveLight(id string, stamp time.Time, light int) {
	point := influxClient.NewPointWithMeasurement("light").
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
func (r *Repo) Values(sensorID string, start, end time.Duration) *dataset.Dataset {
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
