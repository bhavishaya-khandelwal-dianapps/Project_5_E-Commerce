package models

import (
	"time"
)

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null;index" json:"userId"`
	ProductId uint      `gorm:"not null;index" json:"productId"`
	Rating    int       `gorm:"not null" json:"rating"` // 1-5 stars
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User    User    `gorm:"foreignKey:UserId"`
	Product Product `gorm:"foreignKey:ProductId"`
}
