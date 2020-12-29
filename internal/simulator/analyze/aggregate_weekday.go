package analyze

import (
	"time"
)

type SummaryDay struct {
	day time.Weekday

	noOfTrades            int
	noOfProfitTrades      int
	noOfLossTrades        int
	profitTradesInPips    float64
	lossTradesInPips      float64
	netProfitTradesInPips float64
	avgWinPct             float64

	pipsNetProfitGainPct float64
}

func (d SummaryDay) Day() time.Weekday {
	return d.day
}

func (d SummaryDay) NoOfTrades() int {
	return d.noOfTrades
}

func (d SummaryDay) NoOfProfitTrades() int {
	return d.noOfProfitTrades
}

func (d SummaryDay) NoOfLossTrades() int {
	return d.noOfLossTrades
}

func (d SummaryDay) ProfitTradesInPips() float64 {
	return d.profitTradesInPips
}

func (d SummaryDay) LossTradesInPips() float64 {
	return d.lossTradesInPips
}

func (d SummaryDay) NetProfitTradesInPips() float64 {
	return d.netProfitTradesInPips
}

func (d SummaryDay) AvgWinPercentage() float64 {
	return d.avgWinPct
}

func (d SummaryDay) PipsNetProfitGainPercentage() float64 {
	return d.pipsNetProfitGainPct
}

func (d *SummaryDay) append(day *Day) {
	d.noOfTrades += day.NoOfTrades
	d.noOfProfitTrades += day.NoOfProfitTrades
	d.noOfLossTrades += day.NoOfLossTrades
	d.profitTradesInPips += day.ProfitTradesInPips
	d.lossTradesInPips += day.LossTradesInPips
	d.avgWinPct = pctI(d.noOfProfitTrades, d.noOfTrades)

	d.netProfitTradesInPips = d.profitTradesInPips + d.lossTradesInPips
}

func (d *SummaryDay) postConstruct(totalNetProfitInPips float64) {
	d.pipsNetProfitGainPct = pct(d.netProfitTradesInPips / totalNetProfitInPips)
}

type SummaryWeekday struct {
	monday    *SummaryDay
	tuesday   *SummaryDay
	wednesday *SummaryDay
	thursday  *SummaryDay
	friday    *SummaryDay
}

func (w SummaryWeekday) Days() []*SummaryDay {
	return []*SummaryDay{w.monday, w.tuesday, w.wednesday, w.thursday, w.friday}
}

func (w SummaryWeekday) postConstruct() {
	totalNetProfitInPips := w.monday.netProfitTradesInPips +
		w.tuesday.netProfitTradesInPips +
		w.wednesday.netProfitTradesInPips +
		w.thursday.netProfitTradesInPips +
		w.friday.netProfitTradesInPips

	w.monday.postConstruct(totalNetProfitInPips)
	w.tuesday.postConstruct(totalNetProfitInPips)
	w.wednesday.postConstruct(totalNetProfitInPips)
	w.thursday.postConstruct(totalNetProfitInPips)
	w.friday.postConstruct(totalNetProfitInPips)
}

func analyzeAggregateWeekday(weekdays []*Weekday) (result *SummaryWeekday) {
	result = newAggregateWeekday()

	for _, r := range weekdays {
		result.monday.append(r.monday)
		result.tuesday.append(r.tuesday)
		result.wednesday.append(r.wednesday)
		result.thursday.append(r.thursday)
		result.friday.append(r.friday)
	}

	result.postConstruct()
	return
}

func newAggregateWeekday() *SummaryWeekday {
	return &SummaryWeekday{
		monday:    newAggregateDay(time.Monday),
		tuesday:   newAggregateDay(time.Tuesday),
		wednesday: newAggregateDay(time.Wednesday),
		thursday:  newAggregateDay(time.Thursday),
		friday:    newAggregateDay(time.Friday),
	}
}

func newAggregateDay(d time.Weekday) *SummaryDay {
	return &SummaryDay{day: d}
}
