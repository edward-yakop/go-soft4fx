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

func weekday(result *analyze.Result) (err error) {
	simFilePath := result.Simulator().FilePath
	outputFilePath := simFilePath + ".byDayOfWeek.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = exportByDayOfWeek(simFilePath, outputFilePath, result, writer)

	return
}

func exportByDayOfWeek(simFilePath string, outFilePath string, result *analyze.Result, writer *csv.Writer) (err error) {
	err = writeByDayOfWeekHeader(writer)
	simFileName := filepath.Base(simFilePath)

	dow := result.DayOfWeek
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

func writeByDayOfWeekHeader(writer *csv.Writer) error {
	return writer.Write(byDayOfWeekHeader())
}

func byDayOfWeekHeader() []string {
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
