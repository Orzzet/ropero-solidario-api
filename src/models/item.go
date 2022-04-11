package models

type Item struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Category uint   `json:"category"`
	Amount   uint   `json:"amount"`
}
