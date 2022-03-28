package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
)

func (s *Service) CreateUser(user models.User) (models.UserOutput, error) {
	userInput, err := user.FormatInput()
	if err != nil {
		return models.UserOutput{}, err
	}
	if result := s.DB.Save(&userInput); result.Error != nil {
		return models.UserOutput{}, result.Error
	}
	newUser, err := s.GetUser(userInput.ID)
	if err != nil {
		return models.UserOutput{}, err
	}
	return newUser, nil
}

func (s *Service) GetUser(ID uint) (models.UserOutput, error) {
	var user models.User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return models.UserOutput{}, result.Error
	}
	return user.FormatOutput(), nil
}

func (s *Service) ApproveUser(ID uint) (models.UserOutput, error) {
	user, err := s.GetUser(ID)
	if err != nil {
		return models.UserOutput{}, err
	}
	user.IsApproved = true
	if result := s.DB.Model(&user).Updates(user); result.Error != nil {
		return models.UserOutput{}, result.Error
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
