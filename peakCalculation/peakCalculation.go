package peakCalculation

import (
	"time"
	"coinmarketcap/model"
	_"sort"
	"sort"
)

//Get the highest price per hour for n hours
func GetHiLowPrice(pastNHours,cointype,tag int) []float64 {
	startTime := time.Now();
	var endTime time.Time
	var Price []float64
	var CurrentPrice float64
	for i :=0; i < pastNHours; i++{
		startTime, endTime = startTime.Add(- time.Hour), startTime
		if 1 == tag {
			CurrentPrice = model.GetHighestPrice(startTime, endTime, cointype)
		} else if 2 == tag {
			CurrentPrice = model.GetLowPrice(startTime, endTime, cointype)
		}
		if (CurrentPrice != 0) {
			Price = append(Price, CurrentPrice)
		}
	}
	sort.Float64s(Price)

	return Price
}