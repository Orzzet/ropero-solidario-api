package models

import (
	"encoding/json"
)

type User struct {
	ID         uint
	Name       string
	Email      string
	Role       string
	IsApproved bool
	Password   string
}

type UserOutput struct {
	ID         uint
	Name       string
	Email      string
	Role       string
	IsApproved bool
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
