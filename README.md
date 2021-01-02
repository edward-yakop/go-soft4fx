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

`<html-file>.weekday.csv`

Contains simple analysis of trades aggregated by days.
Includes order counts, winning percentage, win/loss in pips, net profit in money and percentage for a simulator.

`<html-file>.closeOrders.csv`

Contains parsed closed orders with additional columns:
* SL in pips;
* TP in pips;
* TP/SL ratio;
* Order durations in minutes; and 
* Close in pips

`<html-file>.drawdown.csv`

Use closed orders as baseline with additional columns:
* max drawdown (pips, time);  
* max profit during trades (pips, DD, time); and
* statistics on open/close price
  This is to determine whether the order time was new york close.

`aggregate.closeOrders.csv`

Contains all parsed closed orders for all simulator results.

`aggregate.weekday.csv`

Contains all weekday analysis into a single file.

`summary.weekday.csv`

Contains summary of all simulator's weekday analysis.

`aggregate.drawdown.csv`

Contains all drawdown analysis in a single file

For example:

| Day       	| NoOfTrades 	| NoOfProfitTrades 	| NoOfLossTrades 	| AvgWinPct 	| ProfitTradesInPips 	| LossTradesInPips 	| NetProfitTradesInPips 	| PipsNetProfitGainPct 	|
|-----------	|------------	|------------------	|----------------	|-------------	|--------------------	|------------------	|-----------------------	|----------------------	|
| Thursday  	|   78         	|   67            	|   11             	|   85.90      	|    2837.9            	|     -459.0       	|     2378.9               	|     25.44            	|
| Tuesday   	|   77         	|   67            	|   10             	|   87.01      	|    2792.7            	|     -429.6       	|     2363.1               	|     25.27            	|
| Wednesday 	|   72         	|   60            	|   12             	|   83.33      	|    2463.6            	|     -509.0       	|     1954.6               	|     20.90            	|
| Friday    	|   60         	|   51            	|   9             	|   85.00      	|    1621.4            	|     -191.4       	|     1430.0               	|     15.29            	|
| Monday    	|   48         	|   38            	|   10             	|   79.17      	|    1410.9            	|     -186.6       	|     1224.3               	|     13.09            	|


## Badges
![Go](https://github.com/ed-fx/go-soft4fx/workflows/Go/badge.svg)
