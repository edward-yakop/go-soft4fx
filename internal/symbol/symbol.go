package symbol

import (
	"math"
	"strings"

	"github.com/ed-fx/go-duka/api/instrument"
)

type Symbol struct {
	metadata *instrument.Metadata
	symbol   string
}

func (s *Symbol) Symbol() string {
	return s.symbol
}

func (s Symbol) PointDifference(from float64, to float64) int {
	return int(math.Round((from - to) * s.metadata.DecimalFactor()))
}

func (s Symbol) PipDifference(from float64, to float64) float64 {
	return float64(s.PointDifference(from, to)) / 10
}

const standardLot = float64(100_000)

func (s Symbol) StandardLotPipValueOnBaseCurrency(askPrice float64) float64 {
	return standardLot / askPrice / s.metadata.DecimalFactor() / 10
}

func (s Symbol) PriceToString(v float64) string {
	if math.IsNaN(v) {
		return ""
	}
	return s.metadata.PriceToString(v)
}

func (s Symbol) PricePtrToString(v *float64) string {
	if v == nil {
		return ""
	}
	return s.PriceToString(*v)
}

func IsSymbolSupported(symbol string) bool {
	return instrument.GetMetadata(symbol) != nil
}

func Get(symbol string) (c *Symbol) {
	metadata := instrument.GetMetadata(symbol)
	if metadata != nil {
		return newSymbol(metadata)
	}
	return nil
}

func newSymbol(metadata *instrument.Metadata) *Symbol {
	return &Symbol{
		metadata: metadata,
		symbol:   strings.ToUpper(metadata.Code()),
	}
}
