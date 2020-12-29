package analyze

import (
	"forex/go-soft4fx/internal/simulator"
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

func analyze(simulator *simulator.Simulator) (result *Result, err error) {
	weekday, err := analyzeByWeekday(simulator)
	if err == nil {
		result = &Result{simulator: simulator, weekday: weekday}
	}
	return
}
