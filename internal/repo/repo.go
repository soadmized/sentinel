package repo

import (
	"fmt"
	"sentry/internal/config"
	"time"

	influxClient "github.com/influxdata/influxdb-client-go/v2"
	influxApi "github.com/influxdata/influxdb-client-go/v2/api"
)

type Repo struct {
	lastValues *DataSet
	writer     influxApi.WriteAPI
	reader     influxApi.QueryAPI
}

type DataSet struct {
	Temp      float32
	Light     int
	Motion    bool
	UpdatedAt time.Time
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

func (r *Repo) SaveValues(temp float32, motion bool, light int) error {
	go func() {
		r.saveTemperature(temp)
	}()

	go func() {
		r.saveMotion(motion)
	}()

	go func() {
		r.saveLight(light)
	}()

	return nil
}

func (r *Repo) saveTemperature(temp float32) {
	r.lastValues.Temp = temp

	point := influxClient.NewPointWithMeasurement("temperature").
		AddTag("id", "first"). // TODO sensor ID
		AddField("temperature", temp).
		SetTime(r.lastValues.UpdatedAt) // TODO where do I take stamp to write??

	r.writer.WritePoint(point)
}

func (r *Repo) saveMotion(motion bool) {
	r.lastValues.Motion = motion

	point := influxClient.NewPointWithMeasurement("motion").
		AddTag("id", "first"). // TODO sensor ID
		AddField("motion", motion).
		SetTime(r.lastValues.UpdatedAt) // TODO where do I take stamp to write??

	r.writer.WritePoint(point)
}

func (r *Repo) saveLight(light int) {
	r.lastValues.Light = light

	point := influxClient.NewPointWithMeasurement("light").
		AddTag("id", "first"). // TODO sensor ID
		AddField("light", light).
		SetTime(r.lastValues.UpdatedAt) // TODO where do I take stamp to write??

	r.writer.WritePoint(point)
}

// LastValues gets last stored values
func (r *Repo) LastValues() *DataSet {
	return r.lastValues
}

// Values gets set of values between start and end
func (r *Repo) Values(start, end time.Duration) *DataSet {
	return r.lastValues
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
