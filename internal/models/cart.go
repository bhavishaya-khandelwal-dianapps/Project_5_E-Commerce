package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null;index" json:"userId"`
	ProductId uint      `gorm:"not null;index" json:"productId"`
	Quantity  int       `gorm:"not null; default:1" json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User    User    `gorm:"foreignKey:UserId" json:"user"`
	Product Product `gorm:"foreignKey:ProductId" json:"product"`
}
