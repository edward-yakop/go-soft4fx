package simulator

import (
	"forex/go-soft4fx/internal/symbol"
	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"
	"math"
	"time"
)

type OrderType int

const (
	Buy = iota
	Sell
	BuyLimit
	SellLimit
	BuyStop
	SellStop
	Balance
)

func (d OrderType) String() string {
	return [...]string{
		"Buy",
		"Sell",
		"Buy Limit",
		"Sell Limit",
		"Buy Stop",
		"Sell Stop",
		"Balance",
	}[d]
}

func NewOrderTypeFromString(orderType string) (OrderType, error) {
	switch orderType {
	case "buy":
		return Buy, nil
	case "sell":
		return Sell, nil
	case "Buy Limit":
		return BuyLimit, nil
	case "Sell Limit":
		return SellLimit, nil
	case "Buy Stop":
		return BuyStop, nil
	case "Sell Stop":
		return SellStop, nil
	case "balance":
		return Balance, nil
	}
	return Buy, errors.New("Unrecognized order type [" + orderType + "]")
}

type Order struct {
	Id         int
	OpenTime   time.Time
	Type       OrderType
	Size       *float64
	Symbol     string
	OpenPrice  *float64
	StopLoss   *float64
	TakeProfit *float64
	CloseTime  time.Time
	ClosePrice *float64
	Commission *money.Money
	Taxes      *money.Money
	Swap       *money.Money
	Profit     *money.Money
	Comment    string

	duration     time.Duration
	total        *money.Money
	profitInPips float64
	slInPips     float64
	slInMoney    *money.Money
	tpInPips     float64
	tpInMoney    *money.Money
	tpSLRatio    float64
}

func (o Order) Duration() time.Duration {
	return o.duration
}

func (o *Order) Total() *money.Money {
	return o.total
}

func (o Order) IsWin() bool {
	return o.Total().Amount() >= 0
}

func (o Order) IsLoss() bool {
	return o.Total().Amount() < 0
}

func (o Order) ProfitPips() float64 {
	return o.profitInPips
}

func (o Order) SLPips() float64 {
	return o.slInPips
}

func (o Order) TpPips() float64 {
	return o.tpInPips
}

func (o Order) TpSLRatio() float64 {
	return o.tpSLRatio
}

func (o *Order) PostConstruct() {
	if !o.CloseTime.IsZero() {
		o.duration = o.CloseTime.Sub(o.OpenTime)
	} else {
		o.duration = 0
	}

	o.total = money.New(
		o.Commission.Amount()+
			o.Taxes.Amount()+
			o.Swap.Amount()+
			o.Profit.Amount(),
		o.Profit.Currency().Code,
	)
	o.initTpSL()
}

func (o *Order) initTpSL() {
	var profitInPips = float64(0)
	var slInPips = profitInPips
	var tpInPips = profitInPips
	var tpSlRatio = math.NaN()
	orderSymbol := o.Symbol
	if symbol.IsSymbolSupported(orderSymbol) {
		calc := symbol.Get(orderSymbol)
		openPrice := *o.OpenPrice
		closePrice := *o.ClosePrice

		if o.Type == Buy {
			profitInPips = calc.PipDifference(closePrice, openPrice)
		} else if o.Type == Sell {
			profitInPips = calc.PipDifference(openPrice, closePrice)
		}

		currency := o.Profit.Currency()
		currencyCode := currency.Code
		currencyFractionRatio := math.Pow10(currency.Fraction)
		lotSize := *o.Size
		if o.StopLoss != nil {
			sl := *o.StopLoss
			slInPips = math.Abs(calc.PipDifference(openPrice, sl))
			slInMoneyNormalized := int64(
				calc.StandardLotPipValueOnBaseCurrency(sl) * lotSize * slInPips * currencyFractionRatio,
			)
			o.slInMoney = money.New(
				slInMoneyNormalized*-1,
				currencyCode,
			)
		}
		if o.TakeProfit != nil {
			tp := *o.TakeProfit
			tpInPips = math.Abs(calc.PipDifference(openPrice, tp))
			tpInMoneyNormalized := int64(
				calc.StandardLotPipValueOnBaseCurrency(tp) * lotSize * tpInPips * currencyFractionRatio,
			)
			o.tpInMoney = money.New(
				tpInMoneyNormalized,
				currencyCode,
			)
		}

		if o.slInMoney != nil && o.tpInMoney != nil {
			tpSlRatio = math.Round(tpInPips/slInPips*100) / 100
		}
	}

	o.profitInPips = profitInPips
	o.slInPips = slInPips
	o.tpInPips = tpInPips
	o.tpSLRatio = tpSlRatio
}
