package models

import (
	"time"

	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	PaymentNo   string    `json:"payment_no"`
	PaymentDate time.Time `json:"payment_date"`
	PaymentType string    `json:"type"`
	Total       int64     `json:"total"`
	AccountID   string    `json:"account_id"`
	Status      string    `json:"status"`
}
