package model

import (
	"coinmarketcap/database"
	"coinmarketcap/monitorType"
)

//Get the highest price in the last 48 hours
func GetNumHoursMaxPrice(parameter *monitorType.GetNumHoursMaxPriceParameter) Coinmarketcap {
	var result Coinmarketcap
	db := databaseServer.GetDB()
	db.Where("created_at > ? AND created_at < ? AND cointype = ?", parameter.Past,parameter.Now,parameter.Cointype).Order("profit desc").First(&Coinmarketcap{}).Scan(&result)
	return result
}


func GetNumHoursMinPrice(parameter *monitorType.GetNumHoursMaxPriceParameter) Coinmarketcap {
	var result Coinmarketcap
	db := databaseServer.GetDB()
	db.Where("created_at > ? AND created_at < ? AND cointype = ?", parameter.Past,parameter.Now,parameter.Cointype).Order("profit ASC").First(&Coinmarketcap{}).Scan(&result)
	return result
}
