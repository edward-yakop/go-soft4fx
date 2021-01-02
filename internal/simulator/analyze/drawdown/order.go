package drawdown

import (
	"forex/go-soft4fx/internal/simulator"
	"forex/go-soft4fx/internal/symbol"
	"github.com/ed-fx/go-duka/api/tickdata"
	"math"
	"time"
)

type Order struct {
	order *simulator.Order

	openDiff  float64
	closeDiff float64

	maxDD     float64
	maxDDTime time.Time

	maxProfit     float64
	maxProfitDD   float64
	maxProfitTime time.Time

	skip bool
}

func (o Order) OpenDiff() float64 {
	return o.openDiff
}

func (o Order) CloseDiff() float64 {
	return o.closeDiff
}

func (o Order) MaxDD() float64 {
	return o.maxDD
}

func (o Order) MaxDDTime() time.Time {
	return o.maxDDTime
}

func (o Order) MaxProfit() float64 {
	return o.maxProfit
}

func (o Order) MaxProfitDD() float64 {
	return o.maxProfitDD
}

func (o Order) MaxProfitTime() time.Time {
	return o.maxProfitTime
}

func (o Order) Type() simulator.OrderType {
	return o.order.Type
}

func (o Order) Symbol() string {
	return o.order.Symbol
}

func (o Order) Order() *simulator.Order {
	return o.order
}

func (o Order) isCompleted() bool {
	return o.skip || !math.IsNaN(o.closeDiff)
}

func (o *Order) onTick(s *symbol.Symbol, tickTime time.Time, tick *tickdata.TickData) {
	if o.isCompleted() {
		return
	}

	if !o.handleOpenPrice(s, tickTime, tick) {
		return
	}

	dd := o.calcDDInPips(s, tick)
	o.handleMaxDD(tickTime, dd)
	o.handleMaxProfit(tickTime, dd)

	o.handleClosePrice(s, tickTime, tick)
}

func (o *Order) handleOpenPrice(s *symbol.Symbol, tickTime time.Time, tick *tickdata.TickData) bool {
	timeDiff := tickTime.Sub(o.order.OpenTime)
	if timeDiff < 0 {
		return false
	}

	isOpenNotInit := math.IsNaN(o.openDiff)
	if isOpenNotInit && timeDiff < time.Second {
		openDiff := o.calcOpenDiff(s, tick)

		if openDiff == 0 {
			o.openDiff = 0
			return true
		}

		return false
	}

	if isOpenNotInit {
		o.openDiff = o.calcOpenDiff(s, tick)
	}

	return true
}

func (o Order) calcOpenDiff(s *symbol.Symbol, tick *tickdata.TickData) float64 {
	openPrice := *o.order.OpenPrice

	var openDiff float64
	if o.order.Type == simulator.Buy {
		openDiff = s.PipDifference(tick.Ask, openPrice)
	} else {
		openDiff = s.PipDifference(tick.Bid, openPrice)
	}
	return openDiff
}

func (o Order) calcDDInPips(s *symbol.Symbol, tick *tickdata.TickData) float64 {
	openPrice := *o.order.OpenPrice
	if o.order.Type == simulator.Buy {
		return s.PipDifference(tick.Bid, openPrice)
	}
	return s.PipDifference(openPrice, tick.Ask)
}

func (o *Order) handleMaxDD(tickTime time.Time, dd float64) {
	if math.IsNaN(o.maxDD) {
		o.maxDD = dd
		return
	}

	if dd < o.maxDD {
		o.maxDD = dd
		o.maxDDTime = tickTime
	}
}

// Can only be called *AFTER* handleMaxDD
func (o *Order) handleMaxProfit(time time.Time, dd float64) {
	if math.IsNaN(o.maxProfit) {
		o.maxProfit = dd
		o.maxProfitDD = dd
		o.maxProfitTime = time

		return
	}

	if o.maxProfit < dd {
		o.maxProfit = dd
		o.maxProfitDD = o.maxDD
		o.maxProfitTime = time
	}
}

func (o *Order) handleClosePrice(s *symbol.Symbol, tickTime time.Time, tick *tickdata.TickData) {
	timeDiff := tickTime.Sub(o.order.CloseTime)
	if timeDiff < 0 {
		return
	}

	isNotInit := math.IsNaN(o.closeDiff)
	if isNotInit && timeDiff < time.Second {
		closeDiff := o.calcCloseDiff(s, tick)

		if closeDiff == 0 {
			o.closeDiff = 0
			return
		}

		return
	}

	if isNotInit {
		o.closeDiff = o.calcCloseDiff(s, tick)
	}

	return
}

func (o Order) calcCloseDiff(s *symbol.Symbol, tick *tickdata.TickData) float64 {
	closePrice := *o.order.ClosePrice

	var openDiff float64
	if o.order.Type == simulator.Buy {
		openDiff = s.PipDifference(tick.Bid, closePrice)
	} else {
		openDiff = s.PipDifference(tick.Ask, closePrice)
	}

	return openDiff
}

func (o Order) OpenTime() time.Time {
	return o.order.OpenTime
}

func (o Order) CloseTime() time.Time {
	return o.order.CloseTime
}

func (o Order) Id() int {
	return o.order.Id
}

func newOrder(order *simulator.Order) *Order {
	return &Order{
		order: order,

		openDiff:  math.NaN(),
		closeDiff: math.NaN(),

		maxDD:       math.NaN(),
		maxProfit:   math.NaN(),
		maxProfitDD: math.NaN(),

		skip: !(order.Type == simulator.Buy || order.Type == simulator.Sell),
	}
}
