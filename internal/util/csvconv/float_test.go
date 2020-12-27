package csvconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64With2DecimalExp(t *testing.T) {
	assert.Equal(t, "123.00", Float64With2DecimalExp(123))
	assert.Equal(t, "123.34", Float64With2DecimalExp(123.344))
	assert.Equal(t, "123.35", Float64With2DecimalExp(123.347))
	assert.Equal(t, "-123.35", Float64With2DecimalExp(-123.347))
}

func TestFloat64PtrWith2DecimalExp(t *testing.T) {
	assert.Equal(t, "", Float64PtrWith2DecimalExp(nil))
	testFloat64PtrWith2DecimalExp(t, "123.00", 123)
	testFloat64PtrWith2DecimalExp(t, "123.34", 123.344)
	testFloat64PtrWith2DecimalExp(t, "123.35", 123.347)
	testFloat64PtrWith2DecimalExp(t, "-123.35", -123.347)
}

func testFloat64PtrWith2DecimalExp(t *testing.T, exp string, f float64) bool {
	return assert.Equal(t, exp, Float64PtrWith2DecimalExp(&f))
}
