package monitor

import (
	"coinmarketcap/model"
	"time"
	"coinmarketcap/monitorType"
	"coinmarketcap/sendMail"
)


func MonitorTypePrice(data *model.Coinmarketcap)  {

	parameter := monitorType.GetNumHoursMaxPriceParameter{
		Now : time.Now(),
		Past : time.Now().AddDate(0,0,-2),
		Cointype:data.Cointype,
		Profit:data.Profit,
	}
	resultMax := model.GetNumHoursMaxPrice(&parameter)
	//A drop of 50
	if 50 < resultMax.Profit - parameter.Profit {
		diff := parameter.Profit - parameter.Profit
		sendMail.MailTemplate(diff,data.Cointype)
	}

	resultMin := model.GetNumHoursMinPrice(&parameter)
	if 100 < parameter.Profit - resultMin.Profit {
		diff := parameter.Profit - resultMin.Profit
		sendMail.MailTemplate(diff,data.Cointype)
	}
}

func MonitorTotalPrice(data *model.StatisticalPrice)  {

}