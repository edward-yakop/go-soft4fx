package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/ed-fx/go-soft4fx/internal/simulator"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"strings"
)

func parseDetails(sim *simulator.Simulator, row *goquery.Selection) (err error) {
	if err = validateSectionHeader(row, "Details:"); err != nil {
		return
	}
	details := &simulator.Details{}
	currency := sim.Currency
	row, err = parseDetailsGrossAndNetProfit(currency, details, row.Next().Next())
	if err != nil {
		return
	}
	row, err = parseDetailsProfitFactorAndExpectedPayoff(details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsDrawdown(currency, details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsTotalTradesAndShortBuyCount(details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsProfitAndLossTrades(details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsLargestTrades(currency, details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsAverageTrades(currency, details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsMaximumConsecutive(details, row)
	if err != nil {
		return
	}
	row, err = parseDetailsMaximalConsecutive(currency, details, row)
	if err != nil {
		return
	}

	sim.Details = details
	return
}

func parseDetailsGrossAndNetProfit(currency string, details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	grossProfit, err := parseMoney(currency, cells[1])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail gross profit ["+cells[1]+"]")
		return
	}
	grossLoss, err := parseMoney(currency, cells[3])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail gross loss ["+cells[3]+"]")
		return
	}
	totalNetProfit, err := parseMoney(currency, cells[5])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail total net profit ["+cells[5]+"]")
		return
	}

	details.GrossProfit = grossProfit
	details.GrossProfit = grossLoss
	details.TotalNetProfit = totalNetProfit

	nextRow = row.Next()
	return
}

func parseDetailsProfitFactorAndExpectedPayoff(details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	profitFactor, err := parseFloat(cells[1])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail profit factor ["+cells[1]+"]")
		return
	}
	expectedPayoff, err := parseFloat(cells[3])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail expected payoff ["+cells[3]+"]")
		return
	}

	details.ProfitFactor = profitFactor
	details.ExpectedPayoff = expectedPayoff

	nextRow = row.Next()
	return
}

func parseDetailsDrawdown(currency string, details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	{
		absDD, perr := parseMoney(currency, cells[1])
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail absolute drawdown ["+cells[1]+"]")
			return
		}
		details.AbsoluteDrawdown = absDD
	}

	{
		maximalDDWithPctString := cells[3]
		maxDDBktIndex := strings.Index(maximalDDWithPctString, "(")
		maxDDAmount := maximalDDWithPctString[:(maxDDBktIndex - 1)]
		maxDDMoney, perr := parseMoney(currency, maxDDAmount)
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail maximal drawdown ["+maxDDAmount+"]")
			return
		}
		maxDDPctString := maximalDDWithPctString[(maxDDBktIndex + 1):(len(maximalDDWithPctString) - 2)]
		maxDDPct, perr := parseFloat(maxDDPctString)
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail maximal drawdown pct ["+maxDDPctString+"]")
			return
		}
		details.MaximalDrawdownInMoney = maxDDMoney
		details.MaximalDrawdownInPct = maxDDPct
	}

	{
		relativeDDWithPctString := cells[5]
		relDDBktIndex := strings.Index(relativeDDWithPctString, "(")
		relDDPctString := relativeDDWithPctString[:(relDDBktIndex - 2)]
		relDDPct, perr := parseFloat(relDDPctString)
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail relative drawdown pct ["+relDDPctString+"]")
			return
		}
		relDDAmountString := relativeDDWithPctString[(relDDBktIndex + 1):(len(relativeDDWithPctString) - 1)]
		relDDAmount, perr := parseMoney(currency, relDDAmountString)
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail relative drawdown ["+relDDAmountString+"]")
			return
		}
		details.RelativeDrawdownInPct = relDDPct
		details.RelativeDrawdownInMoney = relDDAmount
	}

	nextRow = row.Next().Next()
	return
}

func parseDetailsTotalTradesAndShortBuyCount(details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	totalTraders, err := strconv.Atoi(cells[1])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail total trades ["+cells[1]+"]")
		return
	}
	details.TotalTrades = totalTraders

	{
		shortPositionsWithPctString := cells[3]
		shortPositionsBktIdx := strings.Index(shortPositionsWithPctString, "(")
		shortPositions, perr := strconv.Atoi(shortPositionsWithPctString[:shortPositionsBktIdx])
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail short positions ["+cells[3]+"]")
			return
		}
		details.ShortPositionsCount = shortPositions
		shortPositionsWonPct, perr := parseFloat(shortPositionsWithPctString[(shortPositionsBktIdx + 1):(len(shortPositionsWithPctString) - 2)])
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail short positions won percent ["+cells[3]+"]")
			return
		}
		details.ShortPositionsWonPct = shortPositionsWonPct
	}

	{
		longPositionsWithPctString := cells[5]
		longPositionsBktIdx := strings.Index(longPositionsWithPctString, "(")
		longPositions, perr := strconv.Atoi(longPositionsWithPctString[:longPositionsBktIdx])
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail long positions ["+cells[5]+"]")
			return
		}
		details.LongPositionsCount = longPositions
		longPositionsWonPct, perr := parseFloat(longPositionsWithPctString[(longPositionsBktIdx + 1):(len(longPositionsWithPctString) - 2)])
		if perr != nil {
			err = errors.Wrap(perr, "Failed to parse detail long positions won percent ["+cells[5]+"]")
			return
		}
		details.LongPositionsWonPct = longPositionsWonPct
	}

	nextRow = row.Next()
	return
}

func parseDetailsProfitAndLossTrades(details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)

	profitTradesWithPctString := cells[1]
	profitTradesBktIdx := strings.Index(profitTradesWithPctString, "(")
	profitTrades, err := strconv.Atoi(profitTradesWithPctString[:profitTradesBktIdx])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail profit trades ["+cells[3]+"]")
		return
	}
	details.ProfitTradesCount = profitTrades
	totalTrades := float64(details.TotalTrades)
	details.ProfitTradesPct = math.Round(float64(profitTrades)/totalTrades*100_00) / 100
	details.LossTradesCount = details.TotalTrades - profitTrades
	details.LossTradesPct = math.Round(float64(details.LossTradesCount)/totalTrades*100_00) / 100

	nextRow = row.Next()
	return
}

func parseDetailsLargestTrades(currency string, details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)

	largestProfitTradeString := cells[2]
	largestProfitTrade, err := parseMoney(currency, largestProfitTradeString)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail largest profit trade ["+cells[2]+"]")
		return
	}
	details.LargestProfitTrade = largestProfitTrade

	largestLossTradeString := cells[4]
	largestLossTrade, err := parseMoney(currency, largestLossTradeString)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail largest loss trade ["+cells[4]+"]")
		return
	}
	details.LargestLossTrade = largestLossTrade

	nextRow = row.Next()
	return
}

func parseDetailsAverageTrades(currency string, details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)

	avgProfitTradeString := cells[2]
	avgProfitTrade, err := parseMoney(currency, avgProfitTradeString)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail average profit trade ["+cells[2]+"]")
		return
	}
	details.AverageProfitTrade = avgProfitTrade

	avgLossTradeString := cells[4]
	avgLossTrade, err := parseMoney(currency, avgLossTradeString)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail average loss trade ["+cells[4]+"]")
		return
	}
	details.AverageLossTrade = avgLossTrade

	nextRow = row.Next()
	return
}

func parseDetailsMaximumConsecutive(details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)

	consecutiveWins, err := strconv.Atoi(cells[2])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse maximum consecutive wins ["+cells[2]+"]")
		return
	}
	details.MaximumConsecutiveWins = consecutiveWins

	consecutiveLosses, err := strconv.Atoi(cells[4])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse maximum consecutive losses ["+cells[4]+"]")
		return
	}
	details.MaximumConsecutiveLosses = consecutiveLosses

	nextRow = row.Next()
	return
}

func parseDetailsMaximalConsecutive(currency string, details *simulator.Details, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)

	maxConsecutiveProfitString := cells[2]
	maxConsecutiveProfit, err := parseMoney(currency, maxConsecutiveProfitString)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail maximal consecutive wins ["+maxConsecutiveProfitString+"]")
		return
	}
	details.MaximalConsecutiveWins = maxConsecutiveProfit

	maxConsecutiveLossString := cells[4]
	maxConsecutiveLoss, err := parseMoney(currency, maxConsecutiveLossString)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse detail maximal consecutive loss ["+maxConsecutiveLossString+"]")
		return
	}
	details.MaximalConsecutiveLoss = maxConsecutiveLoss

	nextRow = row.Next()
	return
}
