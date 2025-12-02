package models

import (
	"time"
)

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserId     uint      `gorm:"not null;index" json:"userId"`
	TotalPrice float64   `gorm:"not null" json:"totalPrice"`
	Status     string    `gorm:"size:50;default:'PENDING'" json:"status"` // PENDING, PAID, SHIPPED, DELIVERED, CANCELLED
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	User User `gorm:"foreignKey:UserId"`
}
