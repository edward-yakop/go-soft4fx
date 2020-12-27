package csvconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDurationExp(t *testing.T) {
	assert.Equal(t, "0.00", DurationExp(0))

	duration := 10*time.Hour + 11*time.Minute + 12*time.Second
	assert.Equal(t, "611.20", DurationExp(duration))
}
