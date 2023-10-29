package models

import (
	"golang.org/x/crypto/bcrypt"
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

func (admin *Admin) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)
	return nil
}
