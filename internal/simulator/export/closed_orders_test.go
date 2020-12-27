package export

import (
	"forex/go-soft4fx/internal/simulator"
	"forex/go-soft4fx/internal/symbol"
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func TestClosedOrderToRow_balance(t *testing.T) {
	open := time.Date(
		2020, 12, 27, 00, 00, 00, 00, time.Local,
	)

	order := &simulator.Order{
		Id:         1,
		OpenTime:   open,
		Type:       simulator.Balance,
		Commission: money.New(0, "GBP"),
		Taxes:      money.New(0, "GBP"),
		Swap:       money.New(0, "GBP"),
		Profit:     money.New(10_000_00, "GBP"),
	}
	order.PostConstruct()
	s := symbol.Get("GBPJPY")
	row := closedOrderToRow("202012.html", s, order)

	assert.Equal(t, 21, len(row))
	assert.Equal(t, "202012.html", row[0])
	assert.Equal(t, "1", row[1])
	assert.Equal(t, "2020.12.27 00:00:00", row[2])
	assert.Equal(t, "Balance", row[3])
	assert.Equal(t, "", row[4])
	assert.Equal(t, "", row[5])
	assert.Equal(t, "", row[6])
	assert.Equal(t, "", row[7])
	assert.Equal(t, "", row[8])
	assert.Equal(t, "0.0", row[9])
	assert.Equal(t, "0.0", row[10])
	assert.Equal(t, "", row[11])
	assert.Equal(t, "", row[12])
	assert.Equal(t, "0.00", row[13])
	assert.Equal(t, "", row[14])
	assert.Equal(t, "0.0", row[15])
	assert.Equal(t, "0.00", row[16])
	assert.Equal(t, "0.00", row[17])
	assert.Equal(t, "0.00", row[18])
	assert.Equal(t, "10000.00", row[19])
	assert.Equal(t, "10000.00", row[20])
}

func TestClosedOrderToRow_nonBalance(t *testing.T) {
	openTime := time.Date(
		2020, 12, 27, 00, 00, 00, 00, time.Local,
	)
	lotSize := 0.67
	openPrice := 142.797
	stopLoss := 142.377
	takeProfit := 143.217
	closeTime := time.Date(
		2020, 12, 28, 01, 00, 00, 00, time.Local,
	)
	closePrice := 142.846

	order := &simulator.Order{
		Id:         1,
		OpenTime:   openTime,
		Type:       simulator.Buy,
		Size:       &lotSize,
		Symbol:     "GBPJPY",
		OpenPrice:  &openPrice,
		StopLoss:   &stopLoss,
		TakeProfit: &takeProfit,
		CloseTime:  closeTime,
		ClosePrice: &closePrice,
		Commission: money.New(-1_00, "GBP"),
		Taxes:      money.New(-2_10, "GBP"),
		Swap:       money.New(-3_20, "GBP"),
		Profit:     money.New(10_000_00, "GBP"),
	}
	order.PostConstruct()
	s := symbol.Get("GBPJPY")
	row := closedOrderToRow("202012.html", s, order)

	assert.Equal(t, 21, len(row))
	assert.Equal(t, "202012.html", row[0])
	assert.Equal(t, "1", row[1])
	assert.Equal(t, "2020.12.27 00:00:00", row[2])
	assert.Equal(t, "Buy", row[3])
	assert.Equal(t, "0.67", row[4])
	assert.Equal(t, "GBPJPY", row[5])
	assert.Equal(t, "142.797", row[6])
	assert.Equal(t, "142.377", row[7])
	assert.Equal(t, "143.217", row[8])
	assert.Equal(t, "42.0", row[9])
	assert.Equal(t, "42.0", row[10])
	assert.Equal(t, "1.00", row[11])
	assert.Equal(t, "2020.12.28 01:00:00", row[12])
	assert.Equal(t, "1500.00", row[13])
	assert.Equal(t, "142.846", row[14])
	assert.Equal(t, "4.9", row[15])
	assert.Equal(t, "-1.00", row[16])
	assert.Equal(t, "-2.10", row[17])
	assert.Equal(t, "-3.20", row[18])
	assert.Equal(t, "10000.00", row[19])
	assert.Equal(t, "9993.70", row[20])
}
