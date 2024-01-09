package api

import (
	"sentinel/internal/model"
)

type Dataset struct {
	Id     string  `json:"id,omitempty"` // unique id of device
	Temp   float32 `json:"temp,omitempty"`
	Light  int     `json:"light,omitempty"`
	Motion bool    `json:"motion,omitempty"`
}

func (d Dataset) ToModel() model.Dataset {
	return model.Dataset{
		Id:     d.Id,
		Temp:   d.Temp,
		Light:  d.Light,
		Motion: d.Motion,
	}
}
