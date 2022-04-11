package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
)

func (s *Service) GetOrders() (orders []models.Order, err error) {
	if result := s.DB.Find(&orders); result.Error != nil {
		err = result.Error
	}
	return
}

func (s *Service) GetOrder() (order models.Order, err error) {
	if result := s.DB.First(&order); result.Error != nil {
		err = result.Error
	}
	return
}

func (s *Service) CreateOrder(data map[string]interface{}) (models.Order, error) {
	order := models.Order{
		Status:         data["status"].(string),
		RequesterName:  data["requesterName"].(string),
		RequesterPhone: data["requesterPhone"].(string),
	}
	for _, item := range data["items"].([]interface{}) {
		itemData := item.(map[string]interface{})
		order.Items = append(order.Items, models.OrderLine{
			ItemID: uint(itemData["id"].(float64)),
			Amount: uint(itemData["amount"].(float64)),
		})
	}
	if result := s.DB.Create(&order); result.Error != nil {
		return models.Order{}, result.Error
	}
	return order, nil
}

func (s *Service) EditOrder(ID uint, data map[string]interface{}) (order models.Order, err error) {
	order = models.Order{
		ID:             ID,
		Status:         data["status"].(string),
		RequesterName:  data["requesterName"].(string),
		RequesterPhone: data["requesterPhone"].(string),
	}
	for _, item := range data["items"].([]interface{}) {
		itemData := item.(map[string]interface{})
		order.Items = append(order.Items, models.OrderLine{
			ItemID: uint(itemData["id"].(float64)),
			Amount: uint(itemData["amount"].(float64)),
		})
	}
	if result := s.DB.Model(&order).Update(&order); result.Error != nil {
		return models.Order{}, result.Error
	}
	return
}
