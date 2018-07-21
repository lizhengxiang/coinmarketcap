package coinmarketcap

import (
	"fmt"
	"coinmarketcap/model"
	"coinmarketcap/monitor"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"coinmarketcap/monitorType"
)

func GetBitcoinPrice (nowPrice,total float64)  float64 {
	return nowPrice * total
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
	getVirtualCurrencyAll,err := GetVirtualCurrencyAll()
	if false == err || 0 == len(getVirtualCurrencyAll){
		return
	}
	SaveVirtualCurrency(getVirtualCurrencyAll)

	for _,v := range  resource {
		price  := getVirtualCurrencyAll[v.Name].Close
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

func SaveVirtualCurrency(getVirtualCurrencyAll map[string]monitorType.Symbols)  {
	for k,v := range getVirtualCurrencyAll {
		model.SaveVirtualCurrency(k,v)
	}

}

var url="https://www.huobi.com/-/x/general/index/constituent_symbol/detail?r=zuvksd9xrq8"

type PriceAll struct {
	Data SymbolsArr `json:"data"`
	Code int	`json:"code"`
}

type SymbolsArr struct{
	SymbolsArray []monitorType.Symbols `json:"symbols"`
}



func GetVirtualCurrencyAll() (map[string]monitorType.Symbols,bool) {
	var PriceAllResult PriceAll
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get price All error")
		return nil,false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return	nil,false
	}
	err = json.Unmarshal(body,&PriceAllResult)
	if err != nil {
		fmt.Println(err)
		return nil,false
	}
	resultMap := make(map[string]monitorType.Symbols)
	for _,v := range PriceAllResult.Data.SymbolsArray {
		key := strings.ToUpper(v.Name)
		resultMap[key] = v
	}
	fmt.Println(resultMap["BTC"].Close)
	return resultMap,true
}