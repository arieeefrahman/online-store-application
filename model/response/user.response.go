package response

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRegisterResponse struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Username  string         `json:"username"`
}
