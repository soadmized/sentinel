package queue

import (
	"time"
)

type record struct {
	ID        string    `json:"id"`
	Temp      float32   `json:"temp"`
	Light     int       `json:"light"`
	Motion    bool      `json:"motion"`
	UpdatedAt time.Time `json:"updated_at"`
}
