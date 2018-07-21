package monitorType

import "time"

type GetNumHoursMaxPriceParameter struct {
	Now time.Time
	Past time.Time
	Cointype int
	Profit float64
}


type Symbols struct {
	Close float64	`json:"close"`
	Name string		`json:"name"`
}