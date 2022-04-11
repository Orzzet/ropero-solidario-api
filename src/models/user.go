package models

import (
	"encoding/json"
)

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Role       string `json:"role"`
	IsApproved bool   `json:"isApproved"`
	Password   string
}

type UserOutput struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	IsApproved bool   `json:"isApproved"`
}

func (u User) MarshalJSON() ([]byte, error) {
	userOutput := UserOutput{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Role:       u.Role,
		IsApproved: u.IsApproved,
	}
	return json.Marshal(&userOutput)
}
