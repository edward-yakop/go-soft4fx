package export

import (
	"encoding/csv"
	"github.com/ed-fx/go-soft4fx/internal/simulator/analyze"
	"github.com/ed-fx/go-soft4fx/internal/simulator/analyze/drawdown"
	"github.com/ed-fx/go-soft4fx/internal/symbol"
	"github.com/ed-fx/go-soft4fx/internal/util/csvconv"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
)

func Drawdown(r *drawdown.Result) (err error) {
	if r == nil || r.Error() != nil {
		return
	}

	outputFilePath := r.Simulator().FilePath + ".drawdown.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writeDrawdownHeader(writer)
	if err == nil {
		err = writeDrawdownOrders(writer, outputFilePath, r)
	}
	return
}

func writeDrawdownOrders(writer *csv.Writer, outputFilePath string, r *drawdown.Result) (err error) {
	simFileName := filepath.Base(r.Simulator().FilePath)
	orders := r.Orders()
	var s = r.Symbol()
	for _, order := range orders {
		row := drawdownOrderToRow(simFileName, s, order)
		err = writer.Write(row)
		if err != nil {
			orderId := strconv.Itoa(order.Id())
			outputFileName := filepath.Base(outputFilePath)
			err = errors.Wrap(
				err,
				"Failed to export dd order for order id ["+orderId+"] to file ["+outputFileName+"]",
			)
			return
		}
	}
	return
}

func writeDrawdownHeader(writer *csv.Writer) error {
	return writer.Write(drawdownHeader())
}

func drawdownHeader() []string {
	return []string{
		"File",
		"Ticket", "Open Time", "Type", "Size", "Symbol", "Open Price", "OpenDiffInPips",
		"SL", "TP", "SLInPips", "TPInPips", "TP/SL Ratio",
		"Close Time", "Duration", "Close Price", "CloseDiff", "CloseInPips",
		"MaxDDInPips", "MaxDDTime", "MaxProfitInPips", "MaxProfitDD", "MaxProfitTime",
		"Commission", "Taxes", "Swap", "Profit", "Total",
	}
}

func drawdownOrderToRow(simFileName string, s *symbol.Symbol, o *drawdown.Order) []string {
	ro := o.Order()
	return []string{
		simFileName,
		strconv.Itoa(ro.Id),
		csvconv.TimeExp(ro.OpenTime),
		ro.Type.String(),
		csvconv.Float64PtrWith2DecimalExp(ro.Size),
		ro.Symbol,
		s.PricePtrToString(ro.OpenPrice),
		csvconv.Float64With1DecimalExp(o.OpenDiff()),
		s.PricePtrToString(ro.StopLoss),
		s.PricePtrToString(ro.TakeProfit),
		csvconv.Float64With1DecimalExp(ro.SLPips()),
		csvconv.Float64With1DecimalExp(ro.TpPips()),
		csvconv.Float64With2DecimalExp(ro.TpSLRatio()),
		csvconv.TimeExp(ro.CloseTime),
		csvconv.DurationExp(ro.Duration()),
		s.PricePtrToString(ro.ClosePrice),
		csvconv.Float64With1DecimalExp(o.CloseDiff()),
		csvconv.Float64With1DecimalExp(ro.ProfitPips()),
		csvconv.Float64With1DecimalExp(o.MaxDD()),
		csvconv.TimeExp(o.MaxDDTime()),
		csvconv.Float64With1DecimalExp(o.MaxProfit()),
		csvconv.Float64With1DecimalExp(o.MaxProfitDD()),
		csvconv.TimeExp(o.MaxProfitTime()),
		csvconv.MoneyAmountExp(ro.Commission),
		csvconv.MoneyAmountExp(ro.Taxes),
		csvconv.MoneyAmountExp(ro.Swap),
		csvconv.MoneyAmountExp(ro.Profit),
		csvconv.MoneyAmountExp(ro.Total()),
	}
}

func ExportAggregateDrawdown(results []*analyze.Result) (err error) {
	outputFilePath := "aggregate.drawdown.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writeDrawdownHeader(writer)
	if err == nil {
		for _, r := range results {
			ddr := r.Drawdown()
			if r == nil || ddr.Error() != nil {
				continue
			}

			err = writeDrawdownOrders(writer, outputFilePath, ddr)
			if err != nil {
				break
			}
		}
	}
	return
}
