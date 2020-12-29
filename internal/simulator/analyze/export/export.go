package export

import (
	"forex/go-soft4fx/internal/simulator/analyze"
	"forex/go-soft4fx/internal/simulator/export"
)

func AggregateResult(result *analyze.AggregateResult) (err error) {
	for _, r := range result.Results() {
		err = Result(r)
		if err != nil {
			return
		}
	}

	err = export.AggregateClosedOrders(result.Simulators())
	if err != nil {
		return
	}

	err = SummaryWeekday(result.Weekday())
	if err != nil {
		return
	}

	err = AggregateWeekday(result.Results())
	return
}

func Result(result *analyze.Result) (err error) {
	err = export.SimClosedOrders(result.Simulator())
	if err == nil {
		err = Weekday(result)
	}
	return
}
