package response

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Price        int            `json:"price"`
	Stock        int            `json:"stock"`
	CategoryName string         `json:"category_name"`
}
