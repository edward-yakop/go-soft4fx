package csvconv

import (
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoneyAmountExp(t *testing.T) {
	assert.Equal(t, "", MoneyAmountExp(nil))
	assert.Equal(t, "0.00", MoneyAmountExp(money.New(0, "GBP")))

	assert.Equal(t, "12345.67", MoneyAmountExp(money.New(12_345_67, "GBP")))
	assert.Equal(t, "-12345.67", MoneyAmountExp(money.New(-12_345_67, "GBP")))

	assert.Equal(t, "12345678.90", MoneyAmountExp(money.New(12_345_678_90, "GBP")))
	assert.Equal(t, "-12345678.90", MoneyAmountExp(money.New(-12_345_678_90, "GBP")))
}
