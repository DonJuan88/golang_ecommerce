package helper

import (
	"github.com/pintarkode/my_api/models"
	"gorm.io/gorm"
)

func CheckItemExists(db *gorm.DB, code string) (bool, error) {
	var item models.Items
	result := db.Where("item_code = ?", code).First(&item)

	// If record found, return true
	if result.RowsAffected > 0 {
		return true, nil
	}

	// If no record found, return false
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	// For other database errors, return false and error
	return false, result.Error
}
