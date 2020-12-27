package csvconv

import (
	"fmt"
	"math"
)

func Float64With2DecimalExp(v float64) string {
	if math.IsNaN(v) {
		return ""
	}

	return fmt.Sprintf("%.2f", v)
}

func Float64With1DecimalExp(v float64) string {
	if math.IsNaN(v) {
		return ""
	}

	return fmt.Sprintf("%.1f", v)
}

func Float64PtrWith2DecimalExp(v *float64) string {
	if v == nil {
		return ""
	}
	return Float64With2DecimalExp(*v)
}
