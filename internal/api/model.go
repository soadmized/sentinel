package api

import (
	"time"

	"github.com/soadmized/sentinel/pkg/dataset"
)

type lastValuesReq struct {
	ID string `json:"id"` // unique id of device
}

type saveValuesReq struct {
	ID     string  `json:"id"`     // unique ID of device
	Temp   float32 `json:"temp"`   // temperature sensor data
	Light  int     `json:"light"`  // light sensor data
	Motion int     `json:"motion"` // motion sensor data
}

func (r saveValuesReq) toModel(stamp time.Time) dataset.Dataset {
	return dataset.Dataset{
		ID:        r.ID,
		Temp:      r.Temp,
		Light:     r.Light,
		Motion:    intToBool(r.Motion),
		UpdatedAt: stamp,
	}
}

func intToBool(n int) bool {
	return n == 1
}
