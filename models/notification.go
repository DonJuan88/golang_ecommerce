package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	NotifId string `json:"notif_id"`
	UserID string `json:"user_id"`
	NotifMessage string `json:"notif_message"`
}