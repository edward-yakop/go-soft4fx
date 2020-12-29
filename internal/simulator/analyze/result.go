package analyze

import "forex/go-soft4fx/internal/simulator"

type Result struct {
	simulator *simulator.Simulator
	weekday   *Weekday
}

func (r Result) Simulator() *simulator.Simulator {
	return r.simulator
}

func (r Result) Weekday() *Weekday {
	return r.weekday
}
