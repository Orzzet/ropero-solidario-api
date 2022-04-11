package models

type Order struct {
	ID             uint        `json:"id" gorm:"primaryKey"`
	Status         string      `json:"status"`
	RequesterName  string      `json:"requesterName"`
	RequesterPhone string      `json:"requesterPhone"`
	Items          []OrderLine `json:"items"`
}

type OrderLine struct {
	ID     uint `json:"lineId" gorm:"primaryKey"`
	ItemID uint `json:"id"`
	Amount uint `json:"amount"`
}
