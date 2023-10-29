package models

import (
	"time"

	"gorm.io/gorm"
)

type LastExecution struct {
	gorm.Model
	LastExec time.Time `json:"last_execution"`
}
