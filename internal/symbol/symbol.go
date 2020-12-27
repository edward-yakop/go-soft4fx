package symbol

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Symbol struct {
	symbol string

	pointDigit int
	pointRatio float64
	pipDigit   int
	pipRatio   float64

	priceFmt string
}

// Generated from ICMarket
var codeToSymbol = map[string]*Symbol{
	"AUDCAD":      newSymbol("AUDCAD", 5),
	"AUDCHF":      newSymbol("AUDCHF", 5),
	"AUDJPY":      newSymbol("AUDJPY", 3),
	"AUDNZD":      newSymbol("AUDNZD", 5),
	"AUDSGD":      newSymbol("AUDSGD", 5),
	"AUDUSD":      newSymbol("AUDUSD", 5),
	"AUS200":      newSymbol("AUS200", 2),
	"BCHUSD":      newSymbol("BCHUSD", 2),
	"BRENT_J0":    newSymbol("BRENT_J0", 2),
	"BRENT_K0":    newSymbol("BRENT_K0", 2),
	"BRENT_M0":    newSymbol("BRENT_M0", 2),
	"BTCUSD":      newSymbol("BTCUSD", 2),
	"CADCHF":      newSymbol("CADCHF", 5),
	"CADJPY":      newSymbol("CADJPY", 3),
	"CHFJPY":      newSymbol("CHFJPY", 3),
	"CHFSGD":      newSymbol("CHFSGD", 5),
	"CHINA50":     newSymbol("CHINA50", 2),
	"COCOA_K0":    newSymbol("COCOA_K0", 0),
	"COFFEE_K0":   newSymbol("COFFEE_K0", 2),
	"CORN_K0":     newSymbol("CORN_K0", 2),
	"COTTON_K0":   newSymbol("COTTON_K0", 2),
	"DE30":        newSymbol("DE30", 2),
	"DSHUSD":      newSymbol("DSHUSD", 2),
	"DXY_H0":      newSymbol("DXY_H0", 3),
	"DXY_M0":      newSymbol("DXY_M0", 3),
	"EMCUSD":      newSymbol("EMCUSD", 4),
	"EOSUSD":      newSymbol("EOSUSD", 4),
	"ES35":        newSymbol("ES35", 2),
	"ETHUSD":      newSymbol("ETHUSD", 2),
	"EURAUD":      newSymbol("EURAUD", 5),
	"EURBOBL_M0":  newSymbol("EURBOBL_M0", 2),
	"EURBUND_M0":  newSymbol("EURBUND_M0", 2),
	"EURCAD":      newSymbol("EURCAD", 5),
	"EURCHF":      newSymbol("EURCHF", 5),
	"EURDKK":      newSymbol("EURDKK", 5),
	"EURGBP":      newSymbol("EURGBP", 5),
	"EURHKD":      newSymbol("EURHKD", 5),
	"EURJPY":      newSymbol("EURJPY", 3),
	"EURNOK":      newSymbol("EURNOK", 5),
	"EURNZD":      newSymbol("EURNZD", 5),
	"EURPLN":      newSymbol("EURPLN", 5),
	"EURSCHA_M0":  newSymbol("EURSCHA_M0", 2),
	"EURSEK":      newSymbol("EURSEK", 5),
	"EURSGD":      newSymbol("EURSGD", 5),
	"EURTRY":      newSymbol("EURTRY", 5),
	"EURUSD":      newSymbol("EURUSD", 5),
	"EURZAR":      newSymbol("EURZAR", 5),
	"F40":         newSymbol("F40", 2),
	"GBPAUD":      newSymbol("GBPAUD", 5),
	"GBPCAD":      newSymbol("GBPCAD", 5),
	"GBPCHF":      newSymbol("GBPCHF", 5),
	"GBPDKK":      newSymbol("GBPDKK", 5),
	"GBPJPY":      newSymbol("GBPJPY", 3),
	"GBPNOK":      newSymbol("GBPNOK", 5),
	"GBPNZD":      newSymbol("GBPNZD", 5),
	"GBPSEK":      newSymbol("GBPSEK", 5),
	"GBPSGD":      newSymbol("GBPSGD", 5),
	"GBPTRY":      newSymbol("GBPTRY", 5),
	"GBPUSD":      newSymbol("GBPUSD", 5),
	"HK50":        newSymbol("HK50", 2),
	"IT40":        newSymbol("IT40", 2),
	"ITBTP10Y_M0": newSymbol("ITBTP10Y_M0", 2),
	"JGB10Y_H0":   newSymbol("JGB10Y_H0", 2),
	"JGB10Y_M0":   newSymbol("JGB10Y_M0", 2),
	"JP225":       newSymbol("JP225", 2),
	"LTCUSD":      newSymbol("LTCUSD", 2),
	"NMCUSD":      newSymbol("NMCUSD", 3),
	"NOKJPY":      newSymbol("NOKJPY", 3),
	"NOKSEK":      newSymbol("NOKSEK", 5),
	"NZDCAD":      newSymbol("NZDCAD", 5),
	"NZDCHF":      newSymbol("NZDCHF", 5),
	"NZDJPY":      newSymbol("NZDJPY", 3),
	"NZDUSD":      newSymbol("NZDUSD", 5),
	"OJ_K0":       newSymbol("OJ_K0", 2),
	"PPCUSD":      newSymbol("PPCUSD", 3),
	"SEKJPY":      newSymbol("SEKJPY", 3),
	"SGDJPY":      newSymbol("SGDJPY", 3),
	"STOXX50":     newSymbol("STOXX50", 2),
	"SOYBEAN_K0":  newSymbol("SOYBEAN_K0", 2),
	"SUGAR_K0":    newSymbol("SUGAR_K0", 2),
	"UK100":       newSymbol("UK100", 2),
	"UKGB_M0":     newSymbol("UKGB_M0", 2),
	"US2000":      newSymbol("US2000", 2),
	"US30":        newSymbol("US30", 2),
	"US500":       newSymbol("US500", 2),
	"USDCAD":      newSymbol("USDCAD", 5),
	"USDCHF":      newSymbol("USDCHF", 5),
	"USDCNH":      newSymbol("USDCNH", 5),
	"USDCZK":      newSymbol("USDCZK", 4),
	"USDDKK":      newSymbol("USDDKK", 5),
	"USDHKD":      newSymbol("USDHKD", 5),
	"USDHUF":      newSymbol("USDHUF", 3),
	"USDJPY":      newSymbol("USDJPY", 3),
	"USDMXN":      newSymbol("USDMXN", 5),
	"USDNOK":      newSymbol("USDNOK", 5),
	"USDPLN":      newSymbol("USDPLN", 5),
	"USDRUB":      newSymbol("USDRUB", 5),
	"USDSEK":      newSymbol("USDSEK", 5),
	"USDSGD":      newSymbol("USDSGD", 5),
	"USDTHB":      newSymbol("USDTHB", 5),
	"USDTRY":      newSymbol("USDTRY", 5),
	"USDZAR":      newSymbol("USDZAR", 5),
	"UST05Y_M0":   newSymbol("UST05Y_M0", 3),
	"UST10Y_M0":   newSymbol("UST10Y_M0", 3),
	"UST30Y_M0":   newSymbol("UST30Y_M0", 2),
	"USTEC":       newSymbol("USTEC", 2),
	"VIX_H0":      newSymbol("VIX_H0", 2),
	"VIX_J0":      newSymbol("VIX_J0", 2),
	"WTI_J0":      newSymbol("WTI_J0", 2),
	"WTI_K0":      newSymbol("WTI_K0", 2),
	"WHEAT_K0":    newSymbol("WHEAT_K0", 2),
	"XAGEUR":      newSymbol("XAGEUR", 3),
	"XAGUSD":      newSymbol("XAGUSD", 3),
	"XAUAUD":      newSymbol("XAUAUD", 2),
	"XAUEUR":      newSymbol("XAUEUR", 2),
	"XAUUSD":      newSymbol("XAUUSD", 2),
	"XBRUSD":      newSymbol("XBRUSD", 2),
	"XNGUSD":      newSymbol("XNGUSD", 4),
	"XPDUSD":      newSymbol("XPDUSD", 2),
	"XPTUSD":      newSymbol("XPTUSD", 2),
	"XRPUSD":      newSymbol("XRPUSD", 4),
	"XTIUSD":      newSymbol("XTIUSD", 2),
}

func newSymbol(symbol string, pointDigit int) *Symbol {
	return &Symbol{
		symbol:     symbol,
		pointDigit: pointDigit,
		pointRatio: math.Pow10(pointDigit),
		pipDigit:   pointDigit - 1,
		pipRatio:   math.Pow10(pointDigit - 1),
		priceFmt:   "%." + strconv.Itoa(pointDigit) + "f",
	}
}

func (s Symbol) Symbol() string {
	return s.symbol
}

func (s Symbol) PointDigit() int {
	return s.pointDigit
}

func (s Symbol) PipDigit() int {
	return s.pointDigit - 1
}

func (s Symbol) PointDifference(from float64, to float64) int {
	return int(math.Round((from - to) * s.pointRatio))
}

func (s Symbol) PipDifference(from float64, to float64) float64 {
	return float64(s.PointDifference(from, to)) / 10
}

const standardLot = float64(100_000)

func (s Symbol) StandardLotPipValueOnBaseCurrency(askPrice float64) float64 {
	return standardLot / askPrice / s.pipRatio
}

func (s Symbol) PriceToString(v float64) string {
	if math.IsNaN(v) {
		return ""
	}
	return fmt.Sprintf(s.priceFmt, v)
}

func (s Symbol) PricePtrToString(v *float64) string {
	if v == nil {
		return ""
	}
	return s.PriceToString(*v)
}

func IsSymbolSupported(symbol string) bool {
	_, ok := codeToSymbol[strings.ToUpper(symbol)]
	return ok
}

func Get(symbol string) (c *Symbol) {
	c, ok := codeToSymbol[strings.ToUpper(symbol)]
	if !ok {
		c = nil
	}
	return
}
