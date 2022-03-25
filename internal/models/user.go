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

type UserService interface {
	CreateUser(user User) (User, error)
	GetUser(ID uint) (User, error)
	GetUsers() ([]User, error)
	ApproveUser(ID uint) (User, error)
	GetUserHashedPassword(email string) (string, error)
}

func (s *Service) CreateUser(user User) (User, error) {
	if result := s.DB.Save(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *Service) GetUser(ID uint) (User, error) {
	var user User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *Service) ApproveUser(ID uint) (User, error) {
	user, err := s.GetUser(ID)
	if err != nil {
		return User{}, err
	}
	user.IsApproved = true
	if result := s.DB.Model(&user).Updates(user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *Service) GetUserHashedPassword(email string) (string, error) {
	var user User
	if result := s.DB.Select("password").Where("email = ?", email).First(&user); result.Error != nil {
		return "", result.Error
	}
	return user.Password, nil
}
