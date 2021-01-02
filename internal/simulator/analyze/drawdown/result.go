package drawdown

import (
	"forex/go-soft4fx/internal/simulator"
	"forex/go-soft4fx/internal/symbol"
	"sort"
)

type Result struct {
	sim *simulator.Simulator
	err error

	symbol *symbol.Symbol
	orders []*Order
}

func (r Result) Simulator() *simulator.Simulator {
	return r.sim
}

func (r Result) Error() error {
	return r.err
}

func (r Result) Symbol() *symbol.Symbol {
	return r.symbol
}

func (r Result) Orders() []*Order {
	return r.orders
}

func (r Result) OrderCount() int {
	return len(r.orders)
}

func newResult(sim *simulator.Simulator) *Result {
	orders, symbolCode := newOrders(sim)
	return &Result{
		sim:    sim,
		symbol: symbol.Get(symbolCode),
		orders: orders,
	}
}

func newOrders(sim *simulator.Simulator) (result []*Order, symbolCode string) {
	rawOrders := sim.ClosedOrders

	// Orders the rawOrders
	sort.Slice(rawOrders, func(i, j int) bool {
		diff := rawOrders[i].OpenTime.Sub(rawOrders[j].OpenTime)
		if diff < 0 {
			return true
		} else if diff == 0 {
			return rawOrders[i].Id < rawOrders[j].Id
		} else {
			return false
		}
	})

	result = make([]*Order, len(rawOrders))
	for i, ro := range rawOrders {
		order := newOrder(ro)

		if !order.isCompleted() && symbolCode == "" {
			symbolCode = order.Symbol()
		}
		result[i] = order
	}
	return
}
