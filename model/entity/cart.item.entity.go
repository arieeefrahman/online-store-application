package entity

type CartItem struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	ProductID  uint    `json:"product_id"`
	CartID     uint    `json:"cart_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice int     `json:"total_price"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID"`
	Cart       Cart    `gorm:"foreignKey:CartID;references:ID`
}
