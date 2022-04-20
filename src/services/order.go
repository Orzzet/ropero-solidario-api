package services

import (
	"github.com/orzzet/ropero-solidario-api/src/models"
)

func (s *Service) GetOrders() (orders []models.Order, err error) {
	if result := s.DB.Preload("Lines").Find(&orders); result.Error != nil {
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
	var orderLines []models.OrderLine
	for _, line := range data["lines"].([]interface{}) {
		itemData := line.(map[string]interface{})
		orderLines = append(orderLines, models.OrderLine{
			ItemID: uint(itemData["itemId"].(float64)),
			Amount: uint(itemData["amount"].(float64)),
		})
	}
	order.Lines = orderLines
	if result := s.DB.Create(&order); result.Error != nil {
		return models.Order{}, result.Error
	}
	return s.GetOrder(order.ID)
}

func (s *Service) GetOrder(ID uint) (order models.Order, err error) {
	if result := s.DB.Preload("Lines").First(&order, ID); result.Error != nil {
		err = result.Error
	}
	return
}

func (s *Service) EditOrder(ID uint, data map[string]interface{}) (order models.Order, err error) {
	order.ID = ID
	if status, ok := data["status"]; ok {
		order.Status = status.(string)
	}
	if requesterName, ok := data["requesterName"]; ok {
		order.RequesterName = requesterName.(string)
	}
	if requesterPhone, ok := data["requesterPhone"]; ok {
		order.RequesterPhone = requesterPhone.(string)
	}
	if lines, ok := data["lines"]; ok {
		var orderLines []models.OrderLine
		for _, line := range lines.([]interface{}) {
			itemData := line.(map[string]interface{})
			orderLines = append(orderLines, models.OrderLine{
				ItemID: uint(itemData["itemId"].(float64)),
				Amount: uint(itemData["amount"].(float64)),
			})
		}
		order.Lines = orderLines
		err = s.deleteOrderLines(ID)
		if err != nil {
			return models.Order{}, err
		}
	}
	if result := s.DB.Model(&order).Update(&order); result.Error != nil {
		return models.Order{}, result.Error
	}
	return s.GetOrder(ID)
}

func (s *Service) deleteOrderLines(ID uint) (err error) {
	if result := s.DB.Delete(&models.OrderLine{}, "order_id = ?", ID); result.Error != nil {
		err = result.Error
	}
	return
}
