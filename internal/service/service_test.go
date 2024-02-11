package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_checkIfOnline(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		lastSeen time.Time
		want     bool
	}{
		{
			name:     "positive",
			lastSeen: time.Now(),
			want:     true,
		},
		{
			name:     "not online",
			lastSeen: time.Time{},
			want:     false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := checkIfOnline(tt.lastSeen)
			assert.Equal(t, tt.want, got)
		})
	}
}
