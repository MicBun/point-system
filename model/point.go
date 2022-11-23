package model

// Point
type Point struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id"`
	Amount int  `json:"amount"`
}
