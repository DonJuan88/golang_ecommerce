package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	PostCode string `json:"postcode"`
}