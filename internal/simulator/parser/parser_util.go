package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func validateSectionHeader(row *goquery.Selection, expSectionHeader string) error {
	sectionHeader := row.Text()
	if sectionHeader != expSectionHeader {
		return errors.New("Expected [" + expSectionHeader + "] and found [" + sectionHeader + "]")
	}
	return nil
}

func rowToStringArray(row *goquery.Selection) []string {
	filtered := row.ChildrenFiltered("td")
	var rowCells = make([]string, len(filtered.Nodes))
	filtered.Each(func(i int, cells *goquery.Selection) {
		rowCells[i] = cells.Text()
	})
	return rowCells
}

var locNewYork *time.Location = nil

func locationNewYork() *time.Location {
	if locNewYork == nil {
		temp, err := time.LoadLocation("EST")
		if err != nil {
			log.Fatal("Unable to load new york location for parsing time")
		}
		locNewYork = temp
	}
	return locNewYork
}

func parseRowDateTime(dateTime string) (time.Time, error) {
	return time.ParseInLocation("2006.01.02 15:04:05", dateTime, locationNewYork())
}

func parseMoney(currency string, amount string) (m *money.Money, err error) {
	moneyCurrency := money.GetCurrency(currency)
	if moneyCurrency == nil {
		err = errors.New("Unrecognized currency [" + currency + "]")
		return
	}
	amountAsFloat, err := parseFloat(amount)
	if err != nil {
		err = errors.New("Not a parsable amount [" + amount + "]")
		return
	}
	amountInt := int64(amountAsFloat * math.Pow10(moneyCurrency.Fraction))
	m = money.New(amountInt, currency)
	return
}

func parseFloat(amount string) (float64, error) {
	return strconv.ParseFloat(removeSpaces(amount), 64)
}

func removeSpaces(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
