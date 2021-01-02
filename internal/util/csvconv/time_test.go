package csvconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeExp(t *testing.T) {
	assert.Equal(t, "", TimeExp(time.Time{}))

	time := time.Date(2020, 12, 30, 12, 40, 33, 3, time.UTC)
	assert.Equal(t, "2020.12.30 12:40:33", TimeExp(time))
}
