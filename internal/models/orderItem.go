package models

import (
	"time"
)

type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderId   uint      `gorm:"not null;index" json:"orderId"`
	ProductId uint      `gorm:"not null;index" json:"productId"`
	Quantity  int       `gorm:"not null;default:1" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"` // Price at the time of order
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Order   Order   `gorm:"foreignKey:OrderId"`
	Product Product `gorm:"foreignKey:ProductId"`
}
