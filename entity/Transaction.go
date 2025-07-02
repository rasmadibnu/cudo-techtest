package entity

import (
	"time"
)

type Transaction struct {
	ID               int               `gorm:"column:id;NOT NULL" json:"id"`
	UserID           int               `gorm:"column:user_id;NOT NULL" json:"user_id"`
	OrderID          string            `gorm:"column:order_id;NOT NULL" json:"order_id"`
	TransactionDate  time.Time         `gorm:"column:transaction_date;NOT NULL" json:"transaction_date"`
	Amount           string            `gorm:"column:amount;NOT NULL" json:"amount"`
	PaymentMethod    string            `gorm:"column:payment_method;NOT NULL" json:"payment_method"`
	DetectionResults []DetectionResult `gorm:"-" json:"detection_results"`
	Status           string            `gorm:"column:status;NOT NULL" json:"status"`
	ParsedAmount     float64           `json:"-"`
	// ConfidenceScore  int               `json:"confidence_score,omitempty"`
	ZScore    float64   `json:"_,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type DetectionResult struct {
	IsSupicious     bool     `json:"is_supicious"`
	ConfidanceScore float64  `json:"confidance_score"`
	Triggers        []string `json:"triggres"`
}

func (m *Transaction) TableName() string {
	return "transactions"
}
