package entity

import "time"

type TotalCost struct {
	Tenant      string
	SpentAmount float64
	Product     string
	Date        time.Time
}
