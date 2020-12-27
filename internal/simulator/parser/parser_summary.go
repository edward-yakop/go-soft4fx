package parser

import (
	"forex/go-soft4fx/internal/simulator"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func parseSummary(sim *simulator.Simulator, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	if err = validateSectionHeader(row, "Summary:"); err != nil {
		return nil, err
	}
	summary := &simulator.Summary{}
	row, err = parseSummaryDepositAndCreditFacility(sim.Currency, summary, row.Next())
	if err != nil {
		return
	}
	row, err = parseSummaryTradesPLAndMargin(sim.Currency, summary, row)
	if err != nil {
		return
	}
	row, err = parseSummaryBalanceEquityAndFreeMargin(sim.Currency, summary, row)
	if err != nil {
		return
	}

	sim.Summary = summary
	nextRow = row.Next() // Skip blank space
	return
}

func parseSummaryDepositAndCreditFacility(currency string, summary *simulator.Summary, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	deposit, err := parseMoney(currency, cells[1])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary deposit ["+cells[1]+"]")
		return
	}
	creditFacility, err := parseMoney(currency, cells[3])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary credit facility ["+cells[3]+"]")
		return
	}

	summary.Deposit = deposit
	summary.CreditFacility = creditFacility

	nextRow = row.Next()
	return
}

func parseSummaryTradesPLAndMargin(currency string, summary *simulator.Summary, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	closedTradesPL, err := parseMoney(currency, cells[1])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary closed trades P/L ["+cells[1]+"]")
		return
	}
	floatingPL, err := parseMoney(currency, cells[3])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary floating P/L ["+cells[3]+"]")
		return
	}
	margin, err := parseMoney(currency, cells[5])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary margin ["+cells[5]+"]")
		return
	}

	summary.ClosedTradePL = closedTradesPL
	summary.FloatingPL = floatingPL
	summary.Margin = margin

	nextRow = row.Next()
	return
}

func parseSummaryBalanceEquityAndFreeMargin(currency string, summary *simulator.Summary, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	cells := rowToStringArray(row)
	balance, err := parseMoney(currency, cells[1])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary balance ["+cells[1]+"]")
		return
	}
	equity, err := parseMoney(currency, cells[3])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary equity ["+cells[3]+"]")
		return
	}
	freeMargin, err := parseMoney(currency, cells[5])
	if err != nil {
		err = errors.Wrap(err, "Failed to parse summary free margin ["+cells[5]+"]")
		return nil, err
	}

	summary.Balance = balance
	summary.Equity = equity
	summary.FreeMargin = freeMargin

	nextRow = row.Next()
	return
}
