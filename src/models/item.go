package models

type Item struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category"`
	Amount     uint   `json:"amount"`
}
