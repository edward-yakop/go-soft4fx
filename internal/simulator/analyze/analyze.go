package analyze

import (
	"github.com/ed-fx/go-soft4fx/internal/simulator"
	"github.com/ed-fx/go-soft4fx/internal/simulator/analyze/drawdown"
	"github.com/pkg/errors"
)

func Analyze(sims []*simulator.Simulator) (result *AggregateResult, err error) {
	result = &AggregateResult{}

	for _, sim := range sims {
		if r, aerr := analyze(sim); aerr != nil {
			return nil, errors.Wrap(aerr, "Fail to analyze ["+sim.FilePath+"]")
		} else {
			result.add(r)
		}
	}

	result.weekday = analyzeAggregateWeekday(result.weekdays())
	return
}

func analyze(sim *simulator.Simulator) (result *Result, err error) {
	result = &Result{simulator: sim}

	weekday, err := analyzeByWeekday(sim)
	result.weekday = weekday

	dd, err := drawdown.Analyze(sim)
	result.dd = dd

	return
}
