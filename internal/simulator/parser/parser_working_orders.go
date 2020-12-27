package parser

import (
	"forex/go-soft4fx/internal/simulator"
	"github.com/PuerkitoBio/goquery"
)

// TODO: Currently not parsing working orders
func parseWorkingOrders(sim *simulator.Simulator, row *goquery.Selection) (*goquery.Selection, error) {
	if err := validateSectionHeader(row, "Working Orders:"); err != nil {
		return nil, err
	}
	row = row.Next()
	for row.Nodes != nil {
		firstCellText := row.ChildrenFiltered("td").First().Text()
		if firstCellText != "Summary:" {
			row = row.Next()
		} else {
			break
		}
	}
	return row, nil
}
