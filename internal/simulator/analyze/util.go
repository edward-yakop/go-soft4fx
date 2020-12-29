package analyze

import "math"

func pctI(dividend int, divisor int) float64 {
	return pct(float64(dividend) / float64(divisor))
}

func pctInt64(dividend int64, divisor int64) float64 {
	return pct(float64(dividend) / float64(divisor))
}

func pct(v float64) float64 {
	return math.Round(v*100_00) / 100
}
