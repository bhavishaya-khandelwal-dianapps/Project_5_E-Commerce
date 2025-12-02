package models

import "time"

type Token struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null;index" json:"userId"`
	Token     string    `gorm:"size:1000;not null;uniqueIndex" json:"token"`
	Type      string    `gorm:"size:100;not null" json:"type"`
	ExpiresAt time.Time `gorm:"not null" json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
