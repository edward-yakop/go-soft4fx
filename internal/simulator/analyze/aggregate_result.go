package analyze

import (
	"github.com/ed-fx/go-soft4fx/internal/simulator"
)

type AggregateResult struct {
	results []*Result

	weekday *SummaryWeekday
}

func (r AggregateResult) Results() []*Result {
	return r.results
}

func (r *AggregateResult) add(result *Result) {
	r.results = append(r.results, result)
}

func (r AggregateResult) Weekday() *SummaryWeekday {
	return r.weekday
}

func (r AggregateResult) weekdays() (result []*Weekday) {
	count := len(r.results)
	result = make([]*Weekday, count)
	for i, r := range r.results {
		result[i] = r.weekday
	}
	return
}

func (r AggregateResult) Simulators() (result []*simulator.Simulator) {
	count := len(r.results)
	result = make([]*simulator.Simulator, count)
	for i, r := range r.results {
		result[i] = r.simulator
	}
	return
}
