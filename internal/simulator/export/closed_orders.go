package export

import (
	"encoding/csv"
	"forex/go-soft4fx/internal/simulator"
	"forex/go-soft4fx/internal/symbol"
	"forex/go-soft4fx/internal/util/csvconv"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
)

func ClosedOrders(simulator *simulator.Simulator) (err error) {
	outputFilePath := simulator.FilePath + ".closeOrders.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = exportClosedOrder(simulator.FilePath, outputFilePath, simulator, writer)

	return
}

func exportClosedOrder(simFilePath string, outputFilePath string, sim *simulator.Simulator, writer *csv.Writer) (err error) {
	err = writeClosedOrderHeader(writer)
	simFileName := filepath.Base(simFilePath)

	orders := sim.ClosedOrders
	var s = getSymbol(orders)
	for _, order := range orders {
		row := closedOrderToRow(simFileName, s, order)
		csvErr := writer.Write(row)
		if csvErr != nil {
			outputFileName := filepath.Base(outputFilePath)
			err = errors.Wrap(
				csvErr,
				"Failed to export closed order for order id ["+strconv.Itoa(order.Id)+"] to file ["+outputFileName+"]",
			)
			return
		}
	}
	return
}

func getSymbol(orders []*simulator.Order) *symbol.Symbol {
	for _, order := range orders {
		if len(order.Symbol) > 0 {
			return symbol.Get(order.Symbol)
		}
	}
	return nil
}

func writeClosedOrderHeader(writer *csv.Writer) error {
	return writer.Write(closedOrderHeader())
}

func closedOrderHeader() []string {
	return []string{
		"File",
		"Ticket", "Open Time", "Type", "Size", "Symbol", "Open Price",
		"SL", "TP", "SLInPips", "TPInPips", "TP/SL Ratio",
		"Close Time", "Duration", "Close Price", "CloseInPips",
		"Commission", "Taxes", "Swap", "Profit", "Total",
	}
}

func closedOrderToRow(simFileName string, s *symbol.Symbol, o *simulator.Order) []string {
	return []string{
		simFileName,
		strconv.Itoa(o.Id),
		csvconv.TimeExp(&o.OpenTime),
		o.Type.String(),
		csvconv.Float64PtrWith2DecimalExp(o.Size),
		o.Symbol,
		s.PricePtrToString(o.OpenPrice),
		s.PricePtrToString(o.StopLoss),
		s.PricePtrToString(o.TakeProfit),
		csvconv.Float64With1DecimalExp(o.SLPips()),
		csvconv.Float64With1DecimalExp(o.TpPips()),
		csvconv.Float64With2DecimalExp(o.TpSLRatio()),
		csvconv.TimeExp(&o.CloseTime),
		csvconv.DurationExp(o.Duration()),
		s.PricePtrToString(o.ClosePrice),
		csvconv.Float64With1DecimalExp(o.ProfitPips()),
		csvconv.MoneyAmountExp(o.Commission),
		csvconv.MoneyAmountExp(o.Taxes),
		csvconv.MoneyAmountExp(o.Swap),
		csvconv.MoneyAmountExp(o.Profit),
		csvconv.MoneyAmountExp(o.Total()),
	}
}
