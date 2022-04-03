package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Email      string
	Role       string
	IsApproved bool
	Password   string
}

type UserInput struct {
	Name     string
	Email    string
	Password string
}

type UserOutput struct {
	gorm.Model
	Name       string
	Email      string
	Role       string
	IsApproved bool
}

func (u User) MarshalJSON() ([]byte, error) {
	userOutput := UserOutput{
		Model:      u.Model,
		Name:       u.Name,
		Email:      u.Email,
		Role:       u.Role,
		IsApproved: u.IsApproved,
	}
	return json.Marshal(&userOutput)
}
