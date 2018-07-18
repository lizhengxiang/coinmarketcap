package model

import (
	"github.com/jinzhu/gorm"
	"coinmarketcap/database"
)

type Monitor struct {
	gorm.Model
	Cointype int
	NowPrice float64
	TriggerNum int
	Types	int
}


func SaveMonitor(data *Monitor)  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&Monitor{})
	db.Create(data)
}

func DeleteMonitor(types,cointype int)  {
	db := databaseServer.GetDB()
	db.Where("Types = ? AND Cointype = ?",types,cointype).Delete(&Monitor{})
}

//Get the highest price in the last 48 hours
func GetMonitorByCointype(cointype,types int) Monitor {
	var result Monitor
	db := databaseServer.GetDB()
	db.Where("cointype = ? AND	 types = ?", cointype,types).First(&Monitor{}).Scan(&result)
	return result
}

func UpdateMonitor(Id uint,TriggerNum int)  {
	db := databaseServer.GetDB()
	db.Model(&Monitor{}).Update(Monitor{TriggerNum:TriggerNum}).Where("id = ?",Id)
}