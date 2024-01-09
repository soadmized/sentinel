package model

import "time"

// Dataset set of value of sensor which device sends
type Dataset struct {
	Id        string    // unique ID of device
	Temp      float32   // temperature sensor data
	Light     int       // light sensor data
	Motion    bool      // motion sensor data
	UpdatedAt time.Time // timestamp of data
}
