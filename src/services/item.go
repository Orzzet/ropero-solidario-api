package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
)

func (s *Service) GetItems() (items []models.Item, err error) {
	if result := s.DB.Find(&items); result.Error != nil {
		err = result.Error
	}
	return items, err
}

func (s *Service) GetItem(ID uint) (item models.Item, err error) {
	item.ID = ID
	if result := s.DB.Find(&item); result.Error != nil {
		err = result.Error
	}
	return item, err
}

func (s *Service) CreateItem(data map[string]interface{}) (models.Item, error) {
	item := models.Item{
		Name:       data["name"].(string),
		CategoryID: uint(data["category"].(float64)),
		Amount:     uint(data["amount"].(float64)),
	}
	if result := s.DB.Create(&item); result.Error != nil {
		return models.Item{}, result.Error
	}
	return s.GetItem(item.ID)
}

func (s *Service) CreateItems(data map[string]interface{}) ([]models.Item, error) {
	itemsData := data["items"].([]interface{})
	for _, item := range itemsData {
		itemData := item.(map[string]interface{})
		item := models.Item{
			Name:       itemData["name"].(string),
			CategoryID: uint(itemData["category"].(float64)),
			Amount:     uint(itemData["amount"].(float64)),
		}
		if result := s.DB.Create(&item); result.Error != nil {
			return []models.Item{}, result.Error
		}
	}
	return s.GetItems()
}

func (s *Service) EditItem(ID uint, data map[string]interface{}) (item models.Item, err error) {
	item = models.Item{
		ID:         ID,
		Name:       data["name"].(string),
		CategoryID: uint(data["category"].(float64)),
		Amount:     uint(data["amount"].(float64)),
	}
	if result := s.DB.Model(&item).Update(&item); result.Error != nil {
		return models.Item{}, result.Error
	}
	return s.GetItem(ID)
}

func (s *Service) DeleteItem(ID uint) error {
	item := models.Item{
		ID: ID,
	}
	if result := s.DB.Delete(&item).Updates(item); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) IsItemInUse(ID uint) bool {
	var orderLine = models.OrderLine{}
	if result := s.DB.Where("item_id = ?", ID).Find(&orderLine); result.Error != nil {
		return false
	}
	return true
}
