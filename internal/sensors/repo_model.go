package sensors

import (
	"time"

	"github.com/uptrace/bun"
)

type model struct {
	bun.BaseModel `bun:"table:sensors"`
	id            int       `bun:"id,pk"`
	sensor_id     string    `bun:"sensor_id"`
	firstSeen     time.Time `bun:"first_seen"`
	lastSeen      time.Time `bun:"last_seen"`
}

type models []model

func (m model) fromModel() Sensor {
	return Sensor{
		ID:        m.sensor_id,
		FirstSeen: m.firstSeen,
		LastSeen:  m.lastSeen,
	}
}

func (m models) fromModels() Sensors {
	res := make(Sensors, 0, len(m))

	for _, mod := range m {
		res = append(res, mod.fromModel())
	}

	return res
}

func toModel(s Sensor) model {
	return model{
		sensor_id: s.ID,
		firstSeen: s.FirstSeen,
		lastSeen:  s.LastSeen,
	}
}

func toModels(s Sensors) models {
	res := make(models, 0, len(s))

	for _, sen := range s {
		res = append(res, toModel(sen))
	}

	return res
}
