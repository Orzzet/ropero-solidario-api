package database

import (
	"github.com/jinzhu/gorm"
	"github.com/orzzet/ropero-solidario-api/src/models"
)

// MigrateDB migrates DB and creates tables
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&models.User{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&models.Category{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&models.Item{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&models.Order{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&models.OrderLine{}); result.Error != nil {
		return result.Error
	}
	return nil
}
