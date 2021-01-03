package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/ed-fx/go-soft4fx/internal/simulator"
)

// TODO: Currently not parsing open orders
func parseOpenTrades(sim *simulator.Simulator, row *goquery.Selection) (*goquery.Selection, error) {
	if err := validateSectionHeader(row, "Open Trades:"); err != nil {
		return nil, err
	}
	row = row.Next()
	for row.Nodes != nil {
		firstCellText := row.ChildrenFiltered("td").First().Text()
		if firstCellText != "Working Orders:" {
			row = row.Next()
		} else {
			break
		}
	}
	return row, nil
}
