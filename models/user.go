package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string     `json:"name"`
	PhoneNumber  string     `json:"phone_number"`
	Email        string     `json:"email"`
	Address      string     `json:"address"`
	MembershipID uint       `json:"membership_id"`
	Membership   Membership `gorm:"foreignKey:MembershipID"`
}

type UserResponse struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	PhoneNumber string             `json:"phone_number"`
	Email       string             `json:"email"`
	Address     string             `json:"address"`
	Membership  MembershipResponse `json:"membership"`
}

func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Address:     u.Address,
		Membership:  u.Membership.ToMembershipResponse(),
	}
}
