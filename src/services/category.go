package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
)

func (s *Service) GetCategories() (categories []models.Category, err error) {
	if result := s.DB.Find(&categories); result.Error != nil {
		err = result.Error
	}
	return categories, err
}

func (s *Service) CreateCategory(data map[string]interface{}) (models.Category, error) {
	category := models.Category{
		Name: data["name"].(string),
	}
	if result := s.DB.Create(&category); result.Error != nil {
		return models.Category{}, result.Error
	}
	return category, nil
}

func (s *Service) CreateCategories(data map[string]interface{}) ([]models.Category, error) {
	categoriesData := data["categories"].([]interface{})
	//categories := make([]models.Category, len(categoriesData))
	for _, category := range categoriesData {
		categoryData := category.(map[string]interface{})
		category := models.Category{
			Name: categoryData["name"].(string),
		}
		if result := s.DB.Create(&category); result.Error != nil {
			return []models.Category{}, result.Error
		}
	}
	return s.GetCategories()
}

func (s *Service) EditCategory(ID uint, data map[string]interface{}) (category models.Category, err error) {
	category = models.Category{
		ID:   ID,
		Name: data["name"].(string),
	}
	if result := s.DB.Model(&category).Update(&category); result.Error != nil {
		return models.Category{}, result.Error
	}
	return
}

func (s *Service) DeleteCategory(ID uint) error {
	category := models.Category{ID: ID}
	if result := s.DB.Delete(&category).Updates(category); result.Error != nil {
		return result.Error
	}
	return nil
}
