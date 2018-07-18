package monitorType

import "time"

type GetNumHoursMaxPriceParameter struct {
	Now time.Time
	Past time.Time
	Cointype int
	Profit float64
}