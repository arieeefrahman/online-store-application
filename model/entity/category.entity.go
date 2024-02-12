package entity

type Category struct {
	ID          string `gorm:"type:char(36);primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
