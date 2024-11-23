package helper

import (
	"github.com/pintarkode/my_api/models"
	"gorm.io/gorm"
)

func CheckCategoryExists(db *gorm.DB, code string) (bool, error) {
	var categories models.Categories
	result := db.Where("category_code = ?", code).First(&categories)

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
