package fstorage

import "sentinel/internal/dataset"

type FastStorage map[string]dataset.Dataset // {sensorID: values}
