package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"size:100;not null" json:"firstName" binding:"required"`
	LastName  string    `gorm:"size:100;not null" json:"lastName" binding:"required"`
	Email     string    `gorm:"size:100;uniqueIndex;" json:"email" binding:"required,email"`
	Password  string    `gorm:"size:255;not null" json:"password" binding:"required,min=6"`
	Role      string    `gorm:"size:50;default:'CUSTOMER'" json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
