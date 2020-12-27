package csvconv

import (
	"fmt"
	"math"
	"time"
)

func DurationExp(d time.Duration) string {
	h := math.Round(d.Minutes()*100) / 100
	return fmt.Sprintf("%.2f", h)
}
