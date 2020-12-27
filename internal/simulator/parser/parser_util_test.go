package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_parseRowDateTime_happyPath(t *testing.T) {
	dateTime, err := parseRowDateTime("2020.01.06 12:21:33")
	assert.NoError(t, err)
	assert.Equal(t, 2020, dateTime.Year())
	assert.Equal(t, time.January, dateTime.Month())
	assert.Equal(t, 6, dateTime.Day())
	assert.Equal(t, 12, dateTime.Hour())
	assert.Equal(t, 21, dateTime.Minute())
	assert.Equal(t, 33, dateTime.Second())
}

func Test_parseMoney_happyPath(t *testing.T) {
	money, err := parseMoney("GBP", "10 000.00")
	assert.NoError(t, err)
	assert.Equal(t, "GBP", money.Currency().Code)
	assert.Equal(t, int64(10_000_00), money.Amount())

	money, err = parseMoney("GBP", "123.45")
	assert.NoError(t, err)
	assert.Equal(t, "GBP", money.Currency().Code)
	assert.Equal(t, int64(123_45), money.Amount())
}
