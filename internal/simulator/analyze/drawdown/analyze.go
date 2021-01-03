package drawdown

import (
	"github.com/ed-fx/go-duka/api/tickdata"
	"github.com/ed-fx/go-duka/api/tickdata/stream"
	"github.com/ed-fx/go-soft4fx/internal/simulator"
	"github.com/ed-fx/go-soft4fx/internal/util"
	"github.com/gookit/gcli/v2/progress"
	"github.com/pkg/errors"
	"log"
	"time"
)

func Analyze(sim *simulator.Simulator) (result *Result, err error) {
	if sim == nil {
		return nil, errors.New("Argument [sim] must not be [null]")
	}

	result = newResult(sim)
	symbol := result.Symbol()
	if symbol == nil {
		return
	}

	orders := result.orders
	log.Println("Analyze drawdown [", sim.FilePath, "]")
	bar := progress.CustomBar(20, progress.BarStyles[0])
	bar.MaxSteps = uint(len(orders))
	bar.Format = progress.MdlBarFormat
	bar.Start()
	for _, o := range orders {
		bar.Advance()
		if o.isCompleted() {
			continue
		}

		stream.
			New(symbol.Symbol(), o.OpenTime(), o.CloseTime().Add(5*time.Minute), util.CacheFolder()).
			EachTick(func(time time.Time, tick *tickdata.TickData, e error) bool {
				if e != nil {
					err = e
					result.err = err
					return false
				}

				o.onTick(symbol, time, tick)
				return !o.isCompleted()
			})
	}
	bar.Finish()

	return
}
