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

func SummaryWeekday(result *analyze.SummaryWeekday) (err error) {
	outputFilePath := "summary.weekday.csv"
	file, err := os.Create(outputFilePath)
	if err != nil {
		return errors.Wrap(err, "Failed to create file ["+outputFilePath+"]")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = summaryWeekdayHeader(writer)
	if err == nil {
		err = summaryWeekdayBody(outputFilePath, result, writer)
	}
	return
}

func summaryWeekdayHeader(writer *csv.Writer) error {
	return writer.Write(summaryWeekdayHeaderContent())
}

func summaryWeekdayHeaderContent() []string {
	return []string{
		"Day",
		"NoOfTrades", "NoOfProfitTrades", "NoOfLossTrades",
		"AvgWinPct",
		"ProfitTradesInPips", "LossTradesInPips", "NetProfitTradesInPips",
		"PipsNetProfitGainPct",
	}
}

func summaryWeekdayBody(outFilePath string, result *analyze.SummaryWeekday, writer *csv.Writer) (err error) {
	for _, day := range result.Days() {
		row := summaryWeekdayToRow(day)
		csvErr := writer.Write(row)
		if csvErr != nil {
			outputFileName := filepath.Base(outFilePath)
			err = errors.Wrap(
				csvErr,
				"Failed to export aggregate weekday ["+day.Day().String()+"] to file ["+outputFileName+"]",
			)
			return
		}
	}
	return
}

func summaryWeekdayToRow(d *analyze.SummaryDay) []string {
	return []string{
		d.Day().String(),
		strconv.Itoa(d.NoOfTrades()),
		strconv.Itoa(d.NoOfProfitTrades()),
		strconv.Itoa(d.NoOfLossTrades()),
		fmt.Sprintf("%.2f", d.AvgWinPercentage()),
		csvconv.Float64With1DecimalExp(d.ProfitTradesInPips()),
		csvconv.Float64With1DecimalExp(d.LossTradesInPips()),
		csvconv.Float64With1DecimalExp(d.NetProfitTradesInPips()),
		csvconv.Float64With2DecimalExp(d.PipsNetProfitGainPercentage()),
	}
}
