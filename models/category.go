package models

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	CategoryCode string `json:"code"`
	CategoryName string `json:"category"`
}
