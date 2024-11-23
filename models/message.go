package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Name string `json:"name"`
	EmailPhone string `json:"email_phone"`
	MyMessage string `json:"message"`
	ReadingStatus bool `json:"status"`
}