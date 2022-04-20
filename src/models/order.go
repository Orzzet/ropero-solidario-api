package models

type Order struct {
	ID             uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	Status         string      `json:"status"`
	RequesterName  string      `json:"requesterName"`
	RequesterPhone string      `json:"requesterPhone"`
	Lines          []OrderLine `json:"lines"`
}

type OrderLine struct {
	ID      uint `json:"lineId" gorm:"primaryKey;autoIncrement"`
	ItemID  uint `json:"itemId"`
	Amount  uint `json:"amount"`
	OrderID uint `json:"orderId"`
}
