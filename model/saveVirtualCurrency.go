package model

import (
	"coinmarketcap/monitorType"
	"github.com/jinzhu/gorm"
	"coinmarketcap/database"
)



type VirtualCurrencyPrice struct {
	gorm.Model
	Price float64
	Name  string
	Cointype int
}


func SaveVirtualCurrency(k string, v monitorType.Symbols)  {
	var SaveVirtualCurrency VirtualCurrencyPrice
		if "BTC" == k {
			SaveVirtualCurrency.Cointype = 1
			SaveVirtualCurrency.Name = "BTC"
			SaveVirtualCurrency.Price = v.Close
			VirtualCurrency(&SaveVirtualCurrency)
		}
		if "ETH" == k {
			SaveVirtualCurrency.Cointype = 2
			SaveVirtualCurrency.Name = "ETH"
			SaveVirtualCurrency.Price = v.Close
			VirtualCurrency(&SaveVirtualCurrency)
		}
		if "EOS" == k {
			SaveVirtualCurrency.Cointype = 3
			SaveVirtualCurrency.Name = "EOS"
			SaveVirtualCurrency.Price = v.Close
			VirtualCurrency(&SaveVirtualCurrency)
		}
}

func VirtualCurrency(data *VirtualCurrencyPrice)  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&VirtualCurrencyPrice{})
	db.Create(data)
}