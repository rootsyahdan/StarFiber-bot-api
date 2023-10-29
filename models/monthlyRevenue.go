package models

import (
	"time"

	"gorm.io/gorm"
)

type MonthlyRevenue struct {
	gorm.Model
	Month   time.Month `json:"Month"`
	Year    int        `json:"Year"`
	Revenue float64    `json:"Revenue"`
}
