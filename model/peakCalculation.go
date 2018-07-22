package model

import (
	"coinmarketcap/database"
	"time"
)

//Get the highest price in the last 48 hours
func GetHighestPrice(startTime,endTime time.Time, Cointype int) float64 {
	var result VirtualCurrencyPrice
	db := databaseServer.GetDB()
	db.Where("created_at > ? AND created_at < ? AND cointype = ?", startTime,endTime,Cointype).
		Order("price desc").First(&result)
	return result.Price
}

