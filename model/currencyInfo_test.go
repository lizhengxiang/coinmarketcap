package model

import (
	"testing"
	"fmt"
)

func TestSaveCurrencyInfo(t *testing.T) {
	testData := CurrencyInfo{
		Name: "BTC",
		BuyingPrice:3410.14,
		CurrencyNum:0.070224900,
		Url:"https://coinmarketcap.com/currencies/bitcoin/",
		Cointype:1,
	}
	SaveCurrencyInfo(&testData)

	testData = CurrencyInfo{
		Name: "ETH",
		BuyingPrice:1122.21,
		CurrencyNum:19.969504300,
		Url:"https://coinmarketcap.com/currencies/ethereum/",
		Cointype:2,
	}
	SaveCurrencyInfo(&testData)

	testData = CurrencyInfo{
		Name: "EOS",
		BuyingPrice:1122.21,
		CurrencyNum:19.969504300,
		Url:"https://coinmarketcap.com/currencies/eos/",
		Cointype:3,
	}
	SaveCurrencyInfo(&testData)

	testData = CurrencyInfo{
		Name: "GNX",
		BuyingPrice:980.60,
		CurrencyNum:998.194908870,
		Url:"https://coinmarketcap.com/currencies/genaro-network/",
		Cointype:4,
	}
	SaveCurrencyInfo(&testData)
}

func TestDeleteCurrencyInfo(t *testing.T) {
	DeleteCurrencyInfo(2)
}

func TestGetCurrencyInfo(t *testing.T)  {
	result := GetCurrencyInfo()
	fmt.Println(result)
	fmt.Println(len(result))
}

func TestGetCurrencyInfoByCointype(t *testing.T)  {
	result := GetCurrencyInfoByCointype(1)
	fmt.Println(result.Name)
}

func TestSaveCurrencyInfoETH(t *testing.T)  {
	DeleteCurrencyInfo(1)

	testData := CurrencyInfo{
		Name: "BTC",
		BuyingPrice:5593.85,
		CurrencyNum:0.11281444,
		Url:"https://coinmarketcap.com/currencies/bitcoin/",
		Cointype:1,
		StartPrice:7442.33,
	}
	SaveCurrencyInfo(&testData)
}