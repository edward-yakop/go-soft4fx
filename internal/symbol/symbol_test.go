package symbol

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIsSymbolSupported_happyPath(t *testing.T) {
	assert.True(t, IsSymbolSupported("gbpjpy"))
	assert.True(t, IsSymbolSupported("GBPJPY"))
	assert.True(t, IsSymbolSupported("eurUsd"))
	assert.True(t, IsSymbolSupported("EURUSD"))
}

func TestIsSymbolSupported_invalidSymbol(t *testing.T) {
	assert.False(t, IsSymbolSupported(""))
	assert.False(t, IsSymbolSupported("GBPJPY_BLA"))
	assert.False(t, IsSymbolSupported("gbpjpy_"))
}

func TestGetPipCalculator_happyPath(t *testing.T) {
	testGetPipCalculator_withValidSymbol(t, "gbpjpy")
	testGetPipCalculator_withValidSymbol(t, "GBPJPY")
	testGetPipCalculator_withValidSymbol(t, "eurUsd")
	testGetPipCalculator_withValidSymbol(t, "EURUSD")
}

func testGetPipCalculator_withValidSymbol(t *testing.T, symbol string) {
	c := Get(symbol)
	assert.NotNil(t, c)
	assert.Equal(t, strings.ToUpper(symbol), c.symbol)
}

func TestGetPipCalculator_invalidSymbol(t *testing.T) {
	assert.Nil(t, Get(""))
	assert.Nil(t, Get("GBPJPY_BLA"))
	assert.Nil(t, Get("gbpjpy_"))
}

func TestPipCalculator_PointDifference(t *testing.T) {
	{
		c := Get("GBPJPY")
		assert.NotNil(t, c)
		assert.Equal(t, -416, c.PointDifference(142.797, 143.213))
	}

	{
		c := Get("EURUSD")
		assert.NotNil(t, c)
		// last 52 weeks range
		assert.Equal(t, 16126, c.PointDifference(1.22683, 1.06557))
	}
}

func TestPipCalculator_PipDifference(t *testing.T) {
	{
		c := Get("GBPJPY")
		assert.NotNil(t, c)
		assert.Equal(t, -41.6, c.PipDifference(142.797, 143.213))
	}

	{
		c := Get("EURUSD")
		assert.NotNil(t, c)
		// last 52 weeks range
		assert.Equal(t, 1612.6, c.PipDifference(1.22683, 1.06557))
	}
}
