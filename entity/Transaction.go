package entity

import (
	"time"
)

type Transactions struct {
	ID              int64     `gorm:"column:id;NOT NULL" json:"id"`
	UserID          int64     `gorm:"column:user_id;NOT NULL" json:"user_id"`
	OrderID         string    `gorm:"column:order_id;NOT NULL" json:"order_id"`
	TransactionDate time.Time `gorm:"column:transaction_date;NOT NULL" json:"transaction_date"`
	Amount          string    `gorm:"column:amount;NOT NULL" json:"amount"`
	PaymentMethod   string    `gorm:"column:payment_method;NOT NULL" json:"payment_method"`
	Status          string    `gorm:"column:status;NOT NULL" json:"status"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *Transactions) TableName() string {
	return "transactions"
}
