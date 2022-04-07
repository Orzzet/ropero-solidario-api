package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(userData map[string]interface{}) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userData["password"].(string)), 8)
	user := models.User{
		Name:     userData["name"].(string),
		Email:    userData["email"].(string),
		Password: string(hashedPassword),
		Role:     userData["role"].(string),
	}
	if result := s.DB.Save(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	newUser, err := s.GetUser(user.ID)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (s *Service) GetUser(ID uint) (models.User, error) {
	var user models.User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *Service) GetUsers() ([]models.User, error) {
	var users []models.User
	if result := s.DB.Find(&users); result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
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

func (s *Service) DeleteUser(ID uint) error {
	user := models.User{
		ID: ID,
	}
	if result := s.DB.Delete(&user).Updates(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) ResetPassword(ID uint, password string) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	user := models.User{
		ID:       ID,
		Password: string(hashedPassword),
	}
	if result := s.DB.Model(&user).Updates(user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
