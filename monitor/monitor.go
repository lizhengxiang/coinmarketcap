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
	Decline(parameter)
	Gain(parameter)
}

func MonitorTotalPrice(data *model.StatisticalPrice)  {

}

func Gain(parameter monitorType.GetNumHoursMaxPriceParameter)  {
	resultMin := model.GetNumHoursMinPrice(&parameter)
	if 100 < parameter.Profit - resultMin.Profit {
		diff := parameter.Profit - resultMin.Profit
		sendMail.MailTemplate(diff,parameter.Cointype)
	}
}

func Decline(parameter monitorType.GetNumHoursMaxPriceParameter)  {
	resultMax := model.GetNumHoursMaxPrice(&parameter)
	//A drop of 20
	GetMonitorByCointype := UpsAndDownsCalculation(parameter.Cointype,1,parameter.Profit)

	if (20 < resultMax.Profit - parameter.Profit && GetMonitorByCointype.Timediff > 20) ||
		(GetMonitorByCointype.TriggerNum > 15) && GetMonitorByCointype.Profitdidd < 0 &&
		GetMonitorByCointype.Timediff > 20 {
			model.DeleteMonitor(1,parameter.Cointype)
			SaveMonitor(parameter.Profit,parameter.Cointype,GetMonitorByCointype.TriggerNum+1, 1)
			diff := parameter.Profit - parameter.Profit
			sendMail.MailTemplate(diff,parameter.Cointype)
	}
}

type UpsAndDowns struct {
	Profitdidd float64
	Timediff 	int64
	TriggerNum  int
}

func UpsAndDownsCalculation(cointype,types int,profit float64)  UpsAndDowns {
	getMonitorByCointype := model.GetMonitorByCointype(cointype,types)
	if 0 == getMonitorByCointype.TriggerNum {
		SaveMonitor(profit,cointype,1,types)
	}
	return UpsAndDowns{
		Profitdidd:profit - getMonitorByCointype.NowPrice,
		TriggerNum: getMonitorByCointype.TriggerNum,
		Timediff:time.Now().Unix() - getMonitorByCointype.CreatedAt.Unix(),
	}
}

func SaveMonitor(nowPrice float64,cointype ,triggerNum ,types  int )  {
	data := model.Monitor {
		Cointype :cointype,
		NowPrice :nowPrice,
		TriggerNum :triggerNum,
		Types:types,
	}
	model.SaveMonitor(&data)
}