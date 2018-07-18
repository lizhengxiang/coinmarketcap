package monitor

import (
	"testing"
	"coinmarketcap/model"
)

func TestMonitorTypePrice(t *testing.T) {
	testData := model.Coinmarketcap{
		Cointype:1,
		Price:1714.7248676147462,
		CoinNum:0.03486,
		Cost:1500,
		Profit:214.72486761474624,
	}
	MonitorTypePrice(&testData)
}
