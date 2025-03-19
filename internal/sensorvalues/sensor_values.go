package sensorvalues

import (
	"time"
)

const (
	Temperature ValueType = "temperature"
	Light       ValueType = "light"
	Motion      ValueType = "motion"
)

type SensorValue struct {
	SensorID  string
	Stamp     time.Time
	Value     float64
	ValueType ValueType
}

type SensorValues []SensorValue

type ValueType string
