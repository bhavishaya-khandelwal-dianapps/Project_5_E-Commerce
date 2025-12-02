package models

import (
	"time"
)

type Payment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderId   uint      `gorm:"not null;index" json:"orderId"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Status    string    `gorm:"size:50;default:'PENDING'" json:"status"` // PENDING, SUCCESS, FAILED
	PaymentId string    `gorm:"size:100" json:"paymentId"`               // Gateway payment ID
	Method    string    `gorm:"size:50" json:"method"`                   // e.g., RAZORPAY, STRIPE
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Order Order `gorm:"foreignKey:OrderId"`
}
