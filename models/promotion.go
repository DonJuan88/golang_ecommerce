package models

import (
	"time"

	"gorm.io/gorm"
)

type Promotions struct {
	gorm.Model
	PromotionCode int       `json:"promotion_code"`
	PromotionName string    `json:"promotion_name"`
	Descriptions  string    `json:"Description"`
	StartingDate  time.Time `json:"starting_date"`
	StartingHour  time.Time `json:"starting_hour"`
	EndingDate    time.Time `json:"ending_date"`
	EndingHour    time.Time `'json:"ending_hour"`
	Active        bool      `json:"active" gorm:"default:false"`
}

type PromotionDetails struct {
	gorm.Model
	PromotionCode int    `json:"promotion_code"`
	ItemCode      string `json:"code"`
	Percentage    int    `json:"percentage"`
}
