package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	NoInvoice       string    `json:"no_invoice"`
	Status          bool      `json:"status"`
	TransactionDate time.Time `json:"transaction_date"`
	UserID          uint      `json:"user_id"`
	User            User      `gorm:"foreignKey:UserID"`
}

type TransactionResponse struct {
	ID              uint         `json:"id"`
	NoInvoice       string       `json:"no_invoice"`
	Status          bool         `json:"status"`
	TransactionDate time.Time    `json:"transaction_date"`
	User            UserResponse `json:"user"`
}

func (t *Transaction) ToTransactionResponse() TransactionResponse {
	return TransactionResponse{
		ID:              t.ID,
		NoInvoice:       t.NoInvoice,
		Status:          t.Status,
		TransactionDate: t.TransactionDate,
		User:            t.User.ToUserResponse(),
	}
}
