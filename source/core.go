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
	resource:=make(map[string]string)
	resource["bitcoin"]="https://coinmarketcap.com/currencies/bitcoin/"
	resource["ethereum"]="https://coinmarketcap.com/currencies/ethereum/"
	resource["eos"]="https://coinmarketcap.com/currencies/eos/"
	resource["gnx"]="https://coinmarketcap.com/currencies/genaro-network/"

	for k,v := range  resource {
		price, err := getPrice(v)
		if nil != err {
			fmt.Println("get price error")
			return
		}
		if "bitcoin" == k {
			Price = GetBitcoinPrice(price,BitcoinTotal)
			Cointype = 1
			CoinNum = BitcoinTotal
			Cost = BitcoinCost
			Profit = Price - Cost
		}
		if "ethereum" == k {
			Price = GetBitcoinPrice(price,EthereumTotal)
			Cointype = 2
			CoinNum = EthereumTotal
			Cost = EthereumCost
			Profit = Price - Cost
		}
		if "eos" == k {
			Price = GetBitcoinPrice(price,EosTotal)
			Cointype = 3
			CoinNum = EosTotal
			Cost = EosCost
			Profit = Price - Cost
		}
		if "gnx" == k {
			Price = GetBitcoinPrice(price,GnxTotal)
			Cointype = 4
			CoinNum = GnxTotal
			Cost = GnxCost
			Profit = Price - Cost
		}
		data := model.Coinmarketcap{
			Cointype:Cointype,
			Price:Price,
			CoinNum:CoinNum,
			Cost:Cost,
			Profit:Profit,
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