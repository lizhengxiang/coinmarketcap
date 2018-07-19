package coinmarketcap

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"coinmarketcap/model"
	"coinmarketcap/monitor"
)

func GetBitcoinPrice (nowPrice,total float64)  float64 {
	return nowPrice * total * USD
}


var Cointype int
var Price float64
var CoinNum float64
var Cost	float64
var Profit	float64

var PriceTotal  float64
var CostTotal 	float64
var ProfitTotal 	float64

func GetAllPrice()  {
	resource := model.GetCurrencyInfo()
	for _,v := range  resource {
		price, err := getPrice(v.Url)
		if nil != err {
			fmt.Println("get price error")
			return
		}
		Price = GetBitcoinPrice(price,v.CurrencyNum)
		Cointype = v.Cointype
		CoinNum = v.CurrencyNum
		Cost = v.BuyingPrice
		Profit = Price - Cost

		data := model.Coinmarketcap{
			Cointype:Cointype,
			Price:Price,
			CoinNum:CoinNum,
			Cost:Cost,
			Profit:Profit,
			BuyTime:v.CreatedAt,
		}
		model.SaveCoinMarketCap(&data)
		monitor.MonitorTypePrice(&data)
		PriceTotal += Price
		CostTotal += Cost
		ProfitTotal += Profit
	}
	statisticalPrice := model.StatisticalPrice{
		PriceTotal:PriceTotal,
		CostTotal:CostTotal,
		ProfitTotal:ProfitTotal,
	}
	model.SaveStatisticalPrice(&statisticalPrice)
	monitor.MonitorTotalPrice(&statisticalPrice)
	//coinmarketcap.SendMail()
	//model.Test()
}

func getPrice(url string) (float64,error){
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}
	var result string
	doc.Find("#quote_price").Each(func(i int, s *goquery.Selection) {
		result = s.Find(".h2").Text()
	})
	return strconv.ParseFloat(result, 32)
}