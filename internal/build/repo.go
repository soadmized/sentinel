package build

import (
	"github.com/soadmized/sentinel/internal/sensors"
	"github.com/soadmized/sentinel/internal/sensorvalues"
)

func (b *Builder) sensorRepo() *sensors.Repo {
	db := b.postgresClient()
	repo := sensors.NewRepo(db)

	return repo
}

func (b *Builder) sensorValuesRepo() *sensorvalues.Repo {
	db := b.postgresClient()
	repo := sensorvalues.NewRepo(db)

	return repo
}
