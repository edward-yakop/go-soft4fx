package csvconv

import (
	"github.com/Rhymond/go-money"
	"strconv"
	"strings"
)

func MoneyAmountExp(m *money.Money) string {
	if m == nil {
		return ""
	}

	amount := m.Amount()

	// The following are copied from money library with thousand and fraction character
	sa := strconv.FormatInt(abs(amount), 10)

	c := m.Currency()
	if len(sa) <= c.Fraction {
		sa = strings.Repeat("0", c.Fraction-len(sa)+1) + sa
	}

	if c.Fraction > 0 {
		sa = sa[:len(sa)-c.Fraction] + "." + sa[len(sa)-c.Fraction:]
	}

	// Add minus sign for negative amount.
	if amount < 0 {
		sa = "-" + sa
	}

	return sa
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
