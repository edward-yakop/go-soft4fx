#!/bin/sh

DIR=`dirname "$0"`
rm *.csv
$DIR/go-soft4fx `pwd`
cat *.byDayOfWeek.csv | awk '!seen[$0]++' > aggregate.byDayOfWeek.csv
cat *.closeOrders.csv | awk '!seen[$0]++' > aggregate.closeOrders.csv
