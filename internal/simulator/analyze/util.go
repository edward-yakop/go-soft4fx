package analyze

import "math"

func pct(v float64) float64 {
	return math.Round(v*100_00) / 100
}
