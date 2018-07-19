package monitor

import (
	"coinmarketcap/model"
	"time"
	"coinmarketcap/monitorType"
	"coinmarketcap/sendMail"
)


func MonitorTypePrice(data *model.Coinmarketcap)  {
	past := time.Now().AddDate(0,0,-1)
	if data.BuyTime.Unix() > time.Now().AddDate(0,0,-2).Unix() {
		past = data.BuyTime
	}
	parameter := monitorType.GetNumHoursMaxPriceParameter{
		Now : time.Now(),
		Past : past,
		Cointype:data.Cointype,
		Profit:data.Profit,
	}
	Decline(parameter)
	Gain(parameter)
}

func MonitorTotalPrice(data *model.StatisticalPrice)  {

}

//现在的价格 - 最近24小时最小值 大于50 && 离第一次报警时间达到30 分钟
//离第一次报警时间达到30 分钟
//当前值-上一次报警的值小于>20 且时间达到30分钟
func Gain(parameter monitorType.GetNumHoursMaxPriceParameter)  {
	resultMin := model.GetNumHoursMinPrice(&parameter)

	GetMonitorByCointype := UpsAndDownsCalculation(parameter.Cointype,2,parameter.Profit)

	if 50 < parameter.Profit - resultMin.Profit && GetMonitorByCointype.Timediff > 1800 ||
		GetMonitorByCointype.TriggerNum > 15 || GetMonitorByCointype.Profitdiff > 20 &&
		GetMonitorByCointype.Timediff > 900{
			model.DeleteMonitor(1,parameter.Cointype)
			SaveMonitor(parameter.Profit,parameter.Cointype,1, 2)
			diff := parameter.Profit - resultMin.Profit
			sendMail.MailTemplate(diff,parameter.Cointype,parameter.Profit)
	}
	if 50 < parameter.Profit - resultMin.Profit {
		UpdateMonitor(GetMonitorByCointype.Id,GetMonitorByCointype.TriggerNum+1)
	}
}

func Decline(parameter monitorType.GetNumHoursMaxPriceParameter)  {
	resultMax := model.GetNumHoursMaxPrice(&parameter)
	//最近24小时最大值-现在的价格大于20 && 离第一次报警时间达到30 分钟
	//离第一次报警时间达到30 分钟
	//当前值-上一次报警的值小于0 且时间达到30分钟
	GetMonitorByCointype := UpsAndDownsCalculation(parameter.Cointype,1,parameter.Profit)
	if 20 < resultMax.Profit - parameter.Profit && GetMonitorByCointype.Timediff > 1800 ||
		GetMonitorByCointype.TriggerNum > 15 || GetMonitorByCointype.Profitdiff < 0 &&
		GetMonitorByCointype.Timediff > 900 {
			model.DeleteMonitor(1,parameter.Cointype)
			SaveMonitor(parameter.Profit,parameter.Cointype,1, 1)
			diff := parameter.Profit - resultMax.Profit
			sendMail.MailTemplate(diff,parameter.Cointype,parameter.Profit)
	}
	if 20 < resultMax.Profit - parameter.Profit {
		UpdateMonitor(GetMonitorByCointype.Id,GetMonitorByCointype.TriggerNum+1)
	}
}

type UpsAndDowns struct {
	Profitdiff float64
	Timediff 	int64
	TriggerNum  int
	Id 	uint
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
		Id:getMonitorByCointype.ID,
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

func UpdateMonitor(Id uint,TriggerNum int)  {
	model.UpdateMonitor(Id, TriggerNum)
}