package models

import (
	"gorm.io/gorm"
)

type Addresses struct {
	gorm.Model
	AccountID    string `json:"accid"`
	AccountTitle string `json:"title"`
	AccountName  string `json:"name"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	Address3     string `json:"address3"`
	PostalCode   string `json:"postalcode"`
	Phone        string `json:"phone"`
	Default      bool   `json:"defa" gorm:"default:true"`
}
