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

	GetMonitorByCointype := UpsAndDownsCalculation(parameter.Cointype,2,parameter.Profit)

	if 50 < parameter.Profit - resultMin.Profit && GetMonitorByCointype.Timediff > 1800 ||
		GetMonitorByCointype.TriggerNum > 15 || GetMonitorByCointype.Profitdiff > 20 &&
		GetMonitorByCointype.Timediff > 1800{
			model.DeleteMonitor(1,parameter.Cointype)
			SaveMonitor(parameter.Profit,parameter.Cointype,GetMonitorByCointype.TriggerNum+1, 2)
			diff := parameter.Profit - resultMin.Profit
			sendMail.MailTemplate(diff,parameter.Cointype)
	}
}

func Decline(parameter monitorType.GetNumHoursMaxPriceParameter)  {
	resultMax := model.GetNumHoursMaxPrice(&parameter)
	GetMonitorByCointype := UpsAndDownsCalculation(parameter.Cointype,1,parameter.Profit)

	if 20 < resultMax.Profit - parameter.Profit && GetMonitorByCointype.Timediff > 1800 ||
		GetMonitorByCointype.TriggerNum > 15 || GetMonitorByCointype.Profitdiff < 0 &&
		GetMonitorByCointype.Timediff > 1800 {
			model.DeleteMonitor(1,parameter.Cointype)
			SaveMonitor(parameter.Profit,parameter.Cointype,GetMonitorByCointype.TriggerNum+1, 1)
			diff := parameter.Profit - resultMax.Profit
			sendMail.MailTemplate(diff,parameter.Cointype)
	}
}

type UpsAndDowns struct {
	Profitdiff float64
	Timediff 	int64
	TriggerNum  int
}

func UpsAndDownsCalculation(cointype,types int,profit float64)  UpsAndDowns {
	getMonitorByCointype := model.GetMonitorByCointype(cointype,types)
	if 0 == getMonitorByCointype.TriggerNum {
		SaveMonitor(profit,cointype,1,types)
		getMonitorByCointype = model.GetMonitorByCointype(cointype,types)
	}
	return UpsAndDowns{
		Profitdiff:profit - getMonitorByCointype.NowPrice,
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