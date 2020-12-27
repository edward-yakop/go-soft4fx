package parser

import (
	s "forex/go-soft4fx/internal/simulator"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

func Parse(htmlFilePath string) (sim *s.Simulator, err error) {
	htmlFile, err := os.Open(htmlFilePath)
	if err != nil {
		err = errors.Wrap(err, "Unable to open html file ["+htmlFilePath+"]")
		return
	}

	defer htmlFile.Close()
	doc, err := goquery.NewDocumentFromReader(htmlFile)
	if err != nil {
		err = errors.Wrap(err, "Failed to open html file ["+htmlFilePath+"]")
		return
	}

	sim = &s.Simulator{
		FilePath: htmlFilePath,
	}
	row := doc.Find("body > div > table > tbody > tr").First()
	if row, err = parseHeader(sim, row); err != nil {
		err = errors.Wrap(err, "Failed to parse file ["+htmlFilePath+"] header")
		return
	}
	if row, err = parseClosedTransactions(sim, row); err != nil {
		err = errors.Wrap(err, "Failed to parse file ["+htmlFilePath+"] closed transactions")
		return
	}
	if row, err = parseOpenTrades(sim, row); err != nil {
		err = errors.Wrap(err, "Failed to parse file ["+htmlFilePath+"] open trades")
		return
	}
	if row, err = parseWorkingOrders(sim, row); err != nil {
		err = errors.Wrap(err, "Failed to parse file ["+htmlFilePath+"] working orders")
		return
	}
	if row, err = parseSummary(sim, row); err != nil {
		err = errors.Wrap(err, "Failed to parse file ["+htmlFilePath+"] summary")
		return
	}
	if err = parseDetails(sim, row); err != nil {
		err = errors.Wrap(err, "Failed to parse file ["+htmlFilePath+"] details")
		return
	}

	sim.PostConstruct()

	return
}

func parseHeader(sim *s.Simulator, row *goquery.Selection) (*goquery.Selection, error) {
	cells := row.ChildrenFiltered("td").First()
	if accountId, err := parseTextWithPrefix(cells, "Account:"); err != nil {
		return nil, errors.Wrap(err, "Account header is not found")
	} else {
		sim.AccountId = accountId
		cells = cells.Next()
	}
	if name, err := parseTextWithPrefix(cells, "Name:"); err != nil {
		return nil, errors.Wrap(err, "Name header is not found")
	} else {
		sim.Name = name
		cells = cells.Next()
	}
	if currency, err := parseTextWithPrefix(cells, "Currency:"); err != nil {
		return nil, errors.Wrap(err, "Currency header is not found")
	} else {
		sim.Currency = currency
		cells = cells.Next()
	}
	if leverage, err := parseTextWithPrefix(cells, "Leverage:"); err != nil {
		return nil, errors.Wrap(err, "Leverage header is not found")
	} else {
		sim.Leverage, _ = strconv.Atoi(leverage[2:])
		cells = cells.Next()
	}

	return row.Next(), nil
}

func parseTextWithPrefix(cells *goquery.Selection, prefix string) (text string, err error) {
	temp := strings.TrimSpace(cells.Text())
	if strings.Index(temp, prefix) == 0 {
		text = temp[len(prefix)+1:]
		err = nil
	} else {
		text = ""
		err = errors.New("[" + prefix + "] prefix is not found in [" + temp + "] string")
	}
	return text, err
}
