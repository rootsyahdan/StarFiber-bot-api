package models

import "gorm.io/gorm"

type Membership struct {
	gorm.Model
	Name        string  `json:"name"`
	Speed       string  `json:"speed"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type MembershipResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Speed       string  `json:"speed"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (m *Membership) ToMembershipResponse() MembershipResponse {
	return MembershipResponse{
		ID:          m.ID,
		Name:        m.Name,
		Speed:       m.Speed,
		Price:       m.Price,
		Description: m.Description,
	}
}
