package models

import (
	"gorm.io/gorm"
)

type Brands struct {
	gorm.Model
	BrandCode string `json:"code"`
	BrandName string `json:"brand"`
}

type CheckBrand struct {
	BrandCode string `json:"code"`
	BrandName string `json:"brand"`
}
