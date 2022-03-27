package user

import "github.com/jinzhu/gorm"

type UserService struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) CreateUser(user User) (User, error) {
	if result := s.DB.Save(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *UserService) GetUser(ID uint) (User, error) {
	var user User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *UserService) ApproveUser(ID uint) (User, error) {
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

func (s *UserService) GetUserHashedPassword(email string) (string, error) {
	var user User
	if result := s.DB.Select("password").Where("email = ?", email).First(&user); result.Error != nil {
		return "", result.Error
	}
	return user.Password, nil
}
