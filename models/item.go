package models

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model         //`json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ItemCode    string `json:"code"`
	Barcode     string `json:"barcode"`
	ItemName    string `json:"name"`
	Description string `json:"desc"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
	BasePrice   int64  `json:"baseprice"`
	SalePrice   int64  `json:"saleprice"`
	Stock       int64  `json:"stock"`
	Unit        string `json:"unit"`
	Active      bool   `json:"active" gorm:"default:true"`
}

type ItemCreteUpdate struct {
	ItemCode    string `gorm:"code"`
	Barcode     string `json:"barcode"`
	ItemName    string `gorm:"name"`
	Description string `gorm:"desc"`
	Category    string `gorm:"category"`
	Brand       string `gorm:"brand"`
	BasePrice   int    `form:"baseprice"`
	SalePrice   int    `form:"saleprice"`
	Stock       int    `form:"stock"`
	Unit        string `gorm:"unit"`
	Active      bool   `json:"active"`
}
