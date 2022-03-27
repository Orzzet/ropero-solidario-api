package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	IsApproved bool   `json:"is_approved"`
	Password   string `json:"password"`
}
