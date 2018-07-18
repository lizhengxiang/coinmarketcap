package model

import (
	"github.com/jinzhu/gorm"
	"coinmarketcap/database"
)

type Coinmarketcap struct {
	gorm.Model
	Cointype int
	Price float64
	CoinNum float64
	Cost	float64
	Profit	float64
}


func SaveCoinMarketCap(data *Coinmarketcap)  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&Coinmarketcap{})
	db.Create(data)
}

func SaveStatisticalPrice(data *StatisticalPrice)  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&StatisticalPrice{})
	db.Create(data)
}

type StatisticalPrice struct {
	gorm.Model
	PriceTotal  float64
	CostTotal 	float64
	ProfitTotal 	float64
}