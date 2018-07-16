package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"coinmarketcap/source"
	"fmt"
)

func getPrice(url string) (float64,error){
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	var result string
	doc.Find("#quote_price").Each(func(i int, s *goquery.Selection) {
		result = s.Find(".h2").Text()
	})
	return strconv.ParseFloat(result, 32)
}

var myBitcoinPrice  float64
var myEthPrice  float64
var myEosPrice  float64
var myGnxPrice  float64

func main() {

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
			myBitcoinPrice = getBitcoinPrice(price,coinmarketcap.BitcoinTotal)
		}
		if "ethereum" == k {
			myEthPrice = getBitcoinPrice(price,coinmarketcap.EthereumTotal)
		}
		if "eos" == k {
			myEosPrice = getBitcoinPrice(price,coinmarketcap.EosTotal)
		}
		if "gnx" == k {
			myGnxPrice = getBitcoinPrice(price,coinmarketcap.GnxTotal)
		}

	}

	fmt.Println("Bitcoin:", myBitcoinPrice)
	fmt.Println("ethereum:", myEthPrice)
	fmt.Println("eos:", myEosPrice)
	fmt.Println("GNX:", myGnxPrice)
	fmt.Println("total",myBitcoinPrice+myEthPrice+myEosPrice+myGnxPrice)
	coinmarketcap.SendMail()
}

func getBitcoinPrice (nowPrice,total float64)  float64 {
	return nowPrice * total * coinmarketcap.USD
}