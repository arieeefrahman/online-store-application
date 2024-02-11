package response

import (
	"time"

	"gorm.io/gorm"
)

type UserRegisterResponse struct {
	ID        string         `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Username  string         `json:"username"`
}
