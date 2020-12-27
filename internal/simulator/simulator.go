package simulator

import (
	"github.com/Rhymond/go-money"
	"math"
)

type Simulator struct {
	FilePath  string
	AccountId string
	Name      string
	Currency  string
	Leverage  int

	ClosedOrders  []*Order
	OpenOrders    []*Order
	WorkingOrders []*Order

	Summary *Summary
	Details *Details

	profitInPips    float64
	lossInPips      float64
	netProfitInPips float64
}

type Summary struct {
	Deposit        *money.Money
	CreditFacility *money.Money
	ClosedTradePL  *money.Money
	FloatingPL     *money.Money
	Margin         *money.Money
	Balance        *money.Money
	Equity         *money.Money
	FreeMargin     *money.Money
}

type Details struct {
	GrossProfit    *money.Money
	GrossLoss      *money.Money
	TotalNetProfit *money.Money

	ProfitFactor   float64
	ExpectedPayoff float64

	AbsoluteDrawdown       *money.Money
	MaximalDrawdownInMoney *money.Money
	MaximalDrawdownInPct   float64

	RelativeDrawdownInMoney *money.Money
	RelativeDrawdownInPct   float64

	TotalTrades          int
	ShortPositionsCount  int
	ShortPositionsWonPct float64
	LongPositionsCount   int
	LongPositionsWonPct  float64

	ProfitTradesCount int
	ProfitTradesPct   float64

	LossTradesCount int
	LossTradesPct   float64

	LargestProfitTrade *money.Money
	LargestLossTrade   *money.Money

	AverageProfitTrade *money.Money
	AverageLossTrade   *money.Money

	MaximumConsecutiveWins   int
	MaximumConsecutiveLosses int

	MaximalConsecutiveWins *money.Money
	MaximalConsecutiveLoss *money.Money
}

func (s *Simulator) PostConstruct() {
	var profitInPips float64 = 0
	var lossInPips float64 = 0
	for _, o := range s.ClosedOrders {
		pips := o.ProfitPips()
		if math.IsNaN(pips) {
			continue
		}

		if o.IsWin() {
			profitInPips += pips
		} else {
			lossInPips += pips
		}
	}

	s.profitInPips = profitInPips
	s.lossInPips = lossInPips
	s.netProfitInPips = profitInPips + lossInPips
}

func (s Simulator) ProfitInPips() float64 {
	return s.profitInPips
}

func (s Simulator) LossInPips() float64 {
	return s.lossInPips
}

func (s Simulator) NetProfitInPips() float64 {
	return s.netProfitInPips
}
