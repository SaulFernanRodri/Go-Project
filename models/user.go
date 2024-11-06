package models

import (
	"time"
)

type User struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	AuthUsername string    `gorm:"type:varchar(100);not null" json:"auth_username"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	BirthDate    time.Time `gorm:"not null" json:"birth_date" binding:"required"`
	Email        string    `gorm:"type:varchar(100);unique;not null" json:"email" binding:"required"`
	Address      string    `gorm:"type:varchar(150)" json:"address"`
	Phone        string    `gorm:"type:varchar(15)" json:"phone" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
