package models

import "gorm.io/gorm"

type ItemImage struct {
	gorm.Model
	ItemCode string `json:"code"`
	FileName string `json:"image"`
}

type ImageUploads struct {
	ItemCode string `json:"code"`
	FileName string `json:"image"`
}
