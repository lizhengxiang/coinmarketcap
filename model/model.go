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

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func GetProvinces()  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&product)
}

func SaveCoinMarketCap(data *Coinmarketcap)  {
	db := databaseServer.GetDB()
	db.AutoMigrate(&Coinmarketcap{})
	db.Create(data)
}