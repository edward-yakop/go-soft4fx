# go-soft4fx

go-soft4fx performs simple analysis of soft4fx html output file.

Currently only orders from "closed transactions" are parsed and analyzed.

## To Install

Download the latest your platform of choice binary file and make that available in the execution PATH.

## To run

### Windows
```
go-soft4fx.exe <directory/html-file>
```

### Linux/MacOS
```
go-soft4fx <directory/html-file>
```

## Output file

`<html-file>.byDayOfWeek.csv`

Contains simple analysis of trades aggregated by days.
Includes order counts, winning percentage, win/loss in pips, net profit in money and percentage for a simulator.

`<html-file>.closeOrders.csv`

Contains parsed orders with a simple helper to quickly determine SL, TP in pips, SL/TP ratio, durations and Close in pips.

## Badges
![Go](https://github.com/ed-fx/go-soft4fx/workflows/Go/badge.svg)
