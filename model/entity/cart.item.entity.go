package entity

import "time"

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"total_price"`
	ProductID uint      `json:"product_id"`
	UserID    string    `json:"user_id" gorm:"type:char(36)"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
}
