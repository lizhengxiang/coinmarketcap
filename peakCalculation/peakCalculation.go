package peakCalculation

import (
	"time"
	"coinmarketcap/model"
	"sort"
)

//Get the highest price per hour for n hours
func GetHighestPrice(pastNHours int) []float64 {
	startTime := time.Now();
	//var resultArr []float64
	var endTime time.Time
	var Price[] float64
	for i :=0; i < pastNHours; i++{
		startTime, endTime = startTime.Add(- time.Hour), startTime
		CurrentPrice := model.GetHighestPrice(startTime, endTime, 1)
		if (CurrentPrice != 0) {
			Price = append(Price, CurrentPrice)
		}
	}
	sort.Float64s(Price)

	return Price
}


