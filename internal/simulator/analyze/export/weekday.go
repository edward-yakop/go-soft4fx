package export

import (
	"encoding/csv"
	"fmt"
	"forex/go-soft4fx/internal/simulator/analyze"
	"forex/go-soft4fx/internal/util/csvconv"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
)

func Weekday(result *analyze.Result) (err error) {
	outputFilePath := result.Simulator().FilePath + ".weekday.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = weekdayHeader(writer)
	if err == nil {
		err = weekdayBody(writer, outputFilePath, result)
	}

	return
}

func weekdayBody(writer *csv.Writer, outFilePath string, result *analyze.Result) (err error) {
	simFileName := filepath.Base(result.Simulator().FilePath)

	dow := result.Weekday()
	for _, day := range dow.Days() {
		row := dayToRow(simFileName, day)
		csvErr := writer.Write(row)
		if csvErr != nil {
			outputFileName := filepath.Base(outFilePath)
			err = errors.Wrap(
				csvErr,
				"Failed to export day of week analysis for day ["+day.Day().String()+"] to file ["+outputFileName+"]",
			)
			return
		}
	}
	return
}

func weekdayHeader(writer *csv.Writer) error {
	return writer.Write(weekdayHeaderContent())
}

func weekdayHeaderContent() []string {
	return []string{
		"File", "Day",
		"NoOfTrades", "NoOfProfitTrades", "NoOfLossTrades", "WinPct",
		"ProfitTradesInPips", "LossTradesInPips",
		"ProfitTradesInMoney", "LossTradesInMoney",
		"ProfitInPipsPct", "LossInPipsPct",
		"NetProfitInMoney", "NetGainInMoneyPct",
	}
}

func dayToRow(simFileName string, d *analyze.Day) []string {
	return []string{
		simFileName,
		d.Day().String(),
		strconv.Itoa(d.NoOfTrades),
		strconv.Itoa(d.NoOfProfitTrades),
		strconv.Itoa(d.NoOfLossTrades),
		fmt.Sprintf("%.2f", d.WinPct()),
		csvconv.Float64With1DecimalExp(d.ProfitTradesInPips),
		csvconv.Float64With1DecimalExp(d.LossTradesInPips),
		csvconv.MoneyAmountExp(d.ProfitTradesInMoney),
		csvconv.MoneyAmountExp(d.LossTradesInMoney),
		csvconv.Float64With1DecimalExp(d.ProfitInPipsPct()),
		csvconv.Float64With1DecimalExp(d.LossInPipsPct()),
		csvconv.MoneyAmountExp(d.NetProfitInMoney()),
		csvconv.Float64With2DecimalExp(d.NetGainInMoneyPct()),
	}
}

func AggregateWeekday(results []*analyze.Result) (err error) {
	outputFilePath := "aggregate.weekday.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = weekdayHeader(writer)
	if err == nil {
		for _, result := range results {
			err = weekdayBody(writer, outputFilePath, result)
		}
	}
	return
}
