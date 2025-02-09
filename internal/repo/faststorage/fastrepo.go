package fstorage

import (
	"github.com/soadmized/sentinel/pkg/dataset"
)

type FastStorage map[string]dataset.Dataset // {sensorID: values}
