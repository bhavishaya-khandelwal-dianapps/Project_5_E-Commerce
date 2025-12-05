package models

import (
	"time"
)

type Payment struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	OrderId           uint      `gorm:"not null;index" json:"orderId"`
	Amount            float64   `gorm:"not null" json:"amount"`
	Status            string    `gorm:"size:50;default:'PENDING'" json:"status"`  // PENDING, SUCCESS, FAILED
	PaymentId         string    `gorm:"size:100" json:"paymentId"`                // Razorpay payment ID
	Method            string    `gorm:"size:50;default:'RAZORPAY'" json:"method"` // Default RAZORPAY
	RazorpayOrderId   string    `gorm:"size:100" json:"razorpayOrderId"`          // Razorpay order ID
	RazorpaySignature string    `gorm:"size:255" json:"razorpaySignature"`        // Razorpay signature
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`

	Order Order `gorm:"foreignKey:OrderId"`
}
