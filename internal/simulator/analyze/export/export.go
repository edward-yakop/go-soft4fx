package export

import (
	"forex/go-soft4fx/internal/simulator/analyze"
	"forex/go-soft4fx/internal/simulator/export"
)

func Export(result *analyze.Result) (err error) {
	err = export.ClosedOrders(result.Simulator())
	if err != nil {
		return
	}
	err = weekday(result)
	if err != nil {

	}
	return
}
