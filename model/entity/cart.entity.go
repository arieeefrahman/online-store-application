package entity

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name       string         `json:"name"`
	Status     string         `json:"status"`
	TotalPrice int            `json:"total_price"`
	UserID     string         `gorm:"type:char(36)"`
	User       User           `gorm:"foreignKey:UserID;references:ID"`
}
