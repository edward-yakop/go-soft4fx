package analyze

import (
	"forex/go-soft4fx/internal/simulator"
	"github.com/pkg/errors"
)

type Result struct {
	simulator *simulator.Simulator
	DayOfWeek *Weekday
}

func (r Result) Simulator() *simulator.Simulator {
	return r.simulator
}

func Analyze(sim *simulator.Simulator) (result *Result, err error) {
	if sim == nil {
		err = errors.New("Argument [sim] must not be [null]")
		return
	}

	dayOfWeek, err := analyzeByDayOfWeek(sim)
	if err != nil {
		return
	}
	result = &Result{
		simulator: sim,
		DayOfWeek: dayOfWeek,
	}

	return
}
