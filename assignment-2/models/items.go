package models

type Item struct {
	ItemId      int    `json:"item_id" gorm:"column:ID; PRIMARY_KEY"`
	ItemCode    int    `json:"item_code" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	OrderId     int
	Order      *Order
}
