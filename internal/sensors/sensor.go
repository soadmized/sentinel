package sensors

import (
	"time"
)

type Sensor struct {
	ID        string
	FirstSeen time.Time
	LastSeen  time.Time
}

type Sensors []Sensor
