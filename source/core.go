package coinmarketcap

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"coinmarketcap/model"
	"coinmarketcap/monitor"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
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
	getVirtualCurrencyAll,err := GetVirtualCurrencyAll()
	if false == err {
		return
	}
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




var url="https://www.huobi.com/-/x/general/index/constituent_symbol/detail?r=zuvksd9xrq8"

type PriceAll struct {
	Data SymbolsArr `json:"data"`
	Code int	`json:"code"`
}

type SymbolsArr struct{
	SymbolsArray []Symbols `json:"symbols"`
}

type Symbols struct {
	Close float64	`json:"close"`
	Name string		`json:"name"`
}

func GetVirtualCurrencyAll() (map[string]Symbols,bool) {
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
	resultMap := make(map[string]Symbols)
	for _,v := range PriceAllResult.Data.SymbolsArray {
		key := strings.ToUpper(v.Name)
		resultMap[key] = v
	}
	fmt.Println(resultMap["BTC"].Close)
	return resultMap,true
}