package sensorvalues

import (
	"time"

	"github.com/uptrace/bun"
)

type model struct {
	bun.BaseModel `bun:"table:sensor_values"`
	id            int       `bun:"id,pk"`
	stamp         time.Time `bun:"stamp"`
	sensorID      string    `bun:"sensor_id"`
	value         float64   `bun:"value"`
	valueType     ValueType `bun:"valueType"`
}

type models []model

func (m model) fromModel() SensorValue {
	return SensorValue{
		SensorID:  m.sensorID,
		Stamp:     m.stamp,
		Value:     m.value,
		ValueType: m.valueType,
	}
}

func (m models) fromModels() SensorValues {
	res := make(SensorValues, 0, len(m))

	for _, mod := range m {
		res = append(res, mod.fromModel())
	}

	return res
}

func toModel(s SensorValue) model {
	return model{
		stamp:     s.Stamp,
		sensorID:  s.SensorID,
		value:     s.Value,
		valueType: s.ValueType,
	}
}

func toModels(s SensorValues) models {
	res := make(models, 0, len(s))

	for _, sen := range s {
		res = append(res, toModel(sen))
	}

	return res
}
