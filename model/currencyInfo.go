package model

import (
	"github.com/jinzhu/gorm"
	"coinmarketcap/database"
)

type CurrencyInfo struct {
	gorm.Model
	Name string
	BuyingPrice float64
	CurrencyNum float64
	Url string
	Cointype int
}


func SaveCurrencyInfo(data *CurrencyInfo)  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&CurrencyInfo{})
	db.Create(data)
}


func DeleteCurrencyInfo(id int)  {
	db := databaseServer.GetDB()
	db.Where("id = ?",id).Delete(&CurrencyInfo{})
}

func GetCurrencyInfo() []CurrencyInfo {
	var CurrencyInfoResults []CurrencyInfo
	db := databaseServer.GetDB()
	db.Find(&CurrencyInfoResults)
	return CurrencyInfoResults
}

func GetCurrencyInfoByCointype(cointype int) CurrencyInfo {
	var CurrencyInfoResults CurrencyInfo
	db := databaseServer.GetDB()
	db.Where("cointype = ?",cointype).Find(&CurrencyInfoResults)
	return CurrencyInfoResults
}

//BTC 0.070224900.000000000.07022490 ≈ 3410.14
//EOS 19.969504300.000000000.02310970 ≈ 1122.21
//GNX 998.194908870.000000000.02019348 ≈ 980.60