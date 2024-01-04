package repo

import "sentry/internal/config"

type Repo struct {
	lastValues *DataSet
}

type DataSet struct {
	temp   float32
	light  int
	motion bool
}

func New(conf *config.Config) (Repo, error) {
	return Repo{}, nil
}

func (r *Repo) SaveTemperature(temp float32) error {
	r.lastValues.temp = temp

	return nil
}

func (r *Repo) SaveMotion(motion bool) error {
	r.lastValues.motion = motion

	return nil
}

func (r *Repo) SaveLight(light int) error {
	r.lastValues.light = light

	return nil
}

func (r *Repo) GetLastValues() *DataSet {
	return r.lastValues
}

func (r *Repo) GetMotion() error {
	return nil
}

func (r *Repo) GetTemperature() error {
	return nil
}
