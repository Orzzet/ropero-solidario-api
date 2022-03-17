package database

import (
	"github.com/jinzhu/gorm"
	"github.com/orzzet/ropero-solidario-api/internal/models"
)

// MigrateDB migrates DB and creates comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&models.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
