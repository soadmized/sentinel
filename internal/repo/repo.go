package repo

import (
	"fmt"
	"sentinel/internal/config"
	"sentinel/internal/model"
	"sync"
	"time"

	influxClient "github.com/influxdata/influxdb-client-go/v2"
	influxApi "github.com/influxdata/influxdb-client-go/v2/api"
)

type Repo struct {
	lastValues model.Dataset // this field is temporary here
	writer     influxApi.WriteAPI
	reader     influxApi.QueryAPI
}

func New(conf config.Config) (Repo, error) {
	url := fmt.Sprintf("http://localhost:%d", conf.Influx.Port)
	client := influxClient.NewClient(url, conf.Influx.Token)

	writer := client.WriteAPI(conf.Influx.Org, conf.Influx.Bucket)
	reader := client.QueryAPI(conf.Influx.Org)

	return Repo{
		writer: writer,
		reader: reader,
	}, nil
}

func (r *Repo) SaveValues(dataset model.Dataset) error {
	r.lastValues.UpdatedAt = dataset.UpdatedAt

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
	r.lastValues.Temp = temp

	point := influxClient.NewPointWithMeasurement("temperature").
		AddTag("id", id).
		AddField("temperature", temp).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

func (r *Repo) saveMotion(id string, stamp time.Time, motion bool) {
	r.lastValues.Motion = motion

	point := influxClient.NewPointWithMeasurement("motion").
		AddTag("id", id).
		AddField("motion", motion).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

func (r *Repo) saveLight(id string, stamp time.Time, light int) {
	r.lastValues.Light = light

	point := influxClient.NewPointWithMeasurement("light").
		AddTag("id", id).
		AddField("light", light).
		SetTime(stamp)

	r.writer.WritePoint(point)
}

// LastValues gets last stored values
func (r *Repo) LastValues() *model.Dataset {
	return &r.lastValues
}

// Values gets set of values between start and end
func (r *Repo) Values(start, end time.Duration) *model.Dataset {
	return &r.lastValues
}

func (r *Repo) getMotion() error {
	return nil
}

func (r *Repo) getTemperature() error {
	return nil
}

func (r *Repo) getLight() error {
	return nil
}
