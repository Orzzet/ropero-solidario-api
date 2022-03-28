package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	gorm.Model
	Name     string `validate:"required"`
	Email    string `validate:"email,required"`
	Password string `validate:"required"`
}

type UserOutput struct {
	gorm.Model
	Name       string
	Email      string
	Role       string
	IsApproved bool
}

func (u *User) FormatInput() (User, error) {
	userInput := UserInput{
		Model:    u.Model,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	if err := validator.New().Struct(userInput); err != nil {
		return User{}, err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), 8)
	return User{
		Model:      userInput.Model,
		Name:       userInput.Name,
		Email:      userInput.Email,
		Password:   string(hashedPassword),
		Role:       "admin",
		IsApproved: false,
	}, nil
}

func (u *User) FormatOutput() UserOutput {
	return UserOutput{
		Model:      u.Model,
		Name:       u.Name,
		Email:      u.Email,
		Role:       u.Role,
		IsApproved: u.IsApproved,
	}
}
