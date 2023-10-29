package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	token    string `json:"token"`
}
