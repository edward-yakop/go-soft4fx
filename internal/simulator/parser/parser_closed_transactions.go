package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Rhymond/go-money"
	"github.com/ed-fx/go-soft4fx/internal/simulator"
	"github.com/pkg/errors"
	"strconv"
)

func parseClosedTransactions(sim *simulator.Simulator, row *goquery.Selection) (*goquery.Selection, error) {
	if err := validateSectionHeader(row, "Closed Transactions:"); err != nil {
		return nil, err
	}
	row = row.Next().Next() // Skip Section and Grid Header
	// Parse InitialDeposit
	row, err := parseInitialDeposit(sim, row)
	if err != nil {
		return nil, err
	}

	var rowCount = 1
	for {
		nextRow, err := parseClosedTransactionOrder(sim, rowCount, row)
		if err != nil {
			return nil, err
		}
		if nextRow == nil {
			break
		}
		rowCount++
		row = nextRow
	}

	for {
		firstTd := row.ChildrenFiltered("td").First().Text()
		if firstTd == "Open Trades:" {
			return row, nil
		}
		row = row.Next()
	}
}

func parseInitialDeposit(sim *simulator.Simulator, row *goquery.Selection) (*goquery.Selection, error) {
	// 0 2020.01.06 00:00:00 balance Initial deposit 10 000.00
	cells := rowToStringArray(row)
	id, err := strconv.Atoi(cells[0])
	if err != nil {
		return nil, errors.New("Failed to parse initial deposit's id of [" + cells[0] + "]")
	}
	openTime, err := parseRowDateTime(cells[1])
	if err != nil {
		return nil, errors.New("Failed to parse initial deposit's openTime of [" + cells[1] + "]")
	}
	orderType, err := simulator.NewOrderTypeFromString(cells[2])
	if err != nil {
		return nil, errors.New("Failed to parse initial deposit's order type of [" + cells[2] + "]")
	}
	amount, err := parseMoney(sim.Currency, cells[4])
	if err != nil {
		return nil, errors.New("Failed to parse initial deposit's amount of [" + cells[4] + "]")
	}

	currency := amount.Currency().Code
	zeroAmount := money.New(0, currency)
	order := &simulator.Order{
		Id:         id,
		OpenTime:   openTime,
		Type:       orderType,
		Commission: zeroAmount,
		Taxes:      zeroAmount,
		Swap:       zeroAmount,
		Profit:     amount,
		Comment:    cells[3],
	}
	order.PostConstruct()
	sim.ClosedOrders = append(sim.ClosedOrders, order)

	return row.Next(), nil
}

func parseClosedTransactionOrder(sim *simulator.Simulator, rowNumber int, row *goquery.Selection) (nextRow *goquery.Selection, err error) {
	rowNumStr := strconv.Itoa(rowNumber)

	//   0 , 1                   , 2   , 3    , 4      , 5       , 6       , 7       , 8                   , 9       , 10   , 11   , 12   , 13
	//	'1','2020.01.13 17:00:00','buy','0.67','gbpjpy','142.797','142.377','143.217','2020.01.13 23:00:00','142.846','0.00','0.00','0.00','22.98'
	cells := rowToStringArray(row)
	id, err := strconv.Atoi(cells[0])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s id of [" + cells[0] + "]")
	}
	openTime, err := parseRowDateTime(cells[1])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s openTime of [" + cells[1] + "]")
	}
	orderType, err := simulator.NewOrderTypeFromString(cells[2])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s order type of [" + cells[2] + "]")
	}
	lotSize, err := parseFloat(cells[3])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s lot size of [" + cells[3] + "]")
	}
	openPrice, err := parseFloat(cells[5])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s open price of [" + cells[5] + "]")
	}
	stopLoss, err := parseFloat(cells[6])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s stop loss of [" + cells[6] + "]")
	}
	takeProfit, err := parseFloat(cells[7])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s take profit of [" + cells[7] + "]")
	}
	closeTime, err := parseRowDateTime(cells[8])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s close time of [" + cells[8] + "]")
	}
	closePrice, err := parseFloat(cells[9])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s close price of [" + cells[9] + "]")
	}
	commission, err := parseMoney(sim.Currency, cells[10])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s commission of [" + cells[10] + "]")
	}
	taxes, err := parseMoney(sim.Currency, cells[11])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s taxes of [" + cells[11] + "]")
	}
	swap, err := parseMoney(sim.Currency, cells[12])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s swap of [" + cells[12] + "]")
	}
	profit, err := parseMoney(sim.Currency, cells[13])
	if err != nil {
		return nil, errors.New("Failed to parse order row [" + rowNumStr + "]'s profit of [" + cells[13] + "]")
	}

	order := &simulator.Order{
		Id:         id,
		OpenTime:   openTime,
		Type:       orderType,
		Size:       &lotSize,
		Symbol:     cells[4],
		OpenPrice:  &openPrice,
		StopLoss:   &stopLoss,
		TakeProfit: &takeProfit,
		CloseTime:  closeTime,
		ClosePrice: &closePrice,
		Commission: commission,
		Taxes:      taxes,
		Swap:       swap,
		Profit:     profit,
	}
	order.PostConstruct()
	sim.ClosedOrders = append(sim.ClosedOrders, order)

	nextRow = closedTransactionNextRow(nextRow, row)
	return
}

func closedTransactionNextRow(nextRow *goquery.Selection, row *goquery.Selection) *goquery.Selection {
	nextRow = row.Next()
	{
		nodes := nextRow.ChildrenFiltered("td").Nodes
		childrenCount := len(nodes)
		if childrenCount == 3 {
			nextRow = nextRow.Next()
		}
	}

	{
		// Just in case if follow with summary
		childrenCount := len(nextRow.ChildrenFiltered("td").Nodes)
		if childrenCount == 5 {
			nextRow = nil
		}
	}
	return nextRow
}
