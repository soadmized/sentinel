package repo

import (
	"fmt"
	"sentinel/internal/config"
	"sentinel/internal/dataset"
	"sync"
	"time"

	influxClient "github.com/influxdata/influxdb-client-go/v2"
	influxApi "github.com/influxdata/influxdb-client-go/v2/api"
)

type fastStorage map[string]dataset.Dataset // {sensorID: values}

type Repo struct {
	fastStorage fastStorage
	writer      influxApi.WriteAPI
	reader      influxApi.QueryAPI
}

func New(conf config.Config) (Repo, error) {
	storage := make(fastStorage)

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
	r.fastStorage[dataset.Id] = dataset

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		r.saveTemperature(dataset.Id, dataset.UpdatedAt, dataset.Temp)

		defer wg.Done()
	}()

	go func() {
		r.saveMotion(dataset.Id, dataset.UpdatedAt, dataset.Motion)

		defer wg.Done()
	}()

	go func() {
		r.saveLight(dataset.Id, dataset.UpdatedAt, dataset.Light)

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

// LastValues gets last stored values
func (r *Repo) LastValues(sensorID string) *dataset.Dataset {
	values, ok := r.fastStorage[sensorID]
	if !ok {
		return &dataset.Dataset{}
	}

	return &values
}

// Values gets set of values between start and end
func (r *Repo) Values(sensorID string, start, end time.Duration) *dataset.Dataset {
	// TODO temporary implementation
	values, ok := r.fastStorage[sensorID]
	if !ok {
		return &dataset.Dataset{}
	}

	return &values
}
