package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
)

func (s *Service) CreateUser(user models.User) (models.User, error) {
	if result := s.DB.Save(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *Service) GetUser(ID uint) (models.User, error) {
	var user models.User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *Service) ApproveUser(ID uint) (models.User, error) {
	user, err := s.GetUser(ID)
	if err != nil {
		return models.User{}, err
	}
	user.IsApproved = true
	if result := s.DB.Model(&user).Updates(user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *Service) GetUserHashedPassword(email string) (string, error) {
	var user models.User
	if result := s.DB.Select("password").Where("email = ?", email).First(&user); result.Error != nil {
		return "", result.Error
	}
	return user.Password, nil
}
