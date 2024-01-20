package api

import (
	"sentinel/internal/dataset"
	"time"
)

type lastValuesReq struct {
	Id string `json:"id"` // unique id of device
}

type saveValuesReq struct {
	Id     string  `json:"id"`     // unique ID of device
	Temp   float32 `json:"temp"`   // temperature sensor data
	Light  int     `json:"light"`  // light sensor data
	Motion bool    `json:"motion"` // motion sensor data
}

func (r saveValuesReq) toModel(stamp time.Time) dataset.Dataset {
	return dataset.Dataset{
		Id:        r.Id,
		Temp:      r.Temp,
		Light:     r.Light,
		Motion:    r.Motion,
		UpdatedAt: stamp,
	}
}
