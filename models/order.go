package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	TransactionNo   string    `json:"transactno"`
	TransactionDate time.Time `json:"transacdate"`
	ItemCode        string    `json:"code"`
	Qty             int64     `json:"qty"`
	BasePrice       int64     `json:"baseprice"`
	SalePrice       int64     `json:"saleprice"`
	ShippingCost    int64     `json:"shippingprice"`
	Discount        int64     `json:"discount"`
	Total           int64     `json:"total"`
	AccountID       string    `json:"accid"`
	PaymentStatus   string    `json:"paymentstatus" gorm:"default:pending"`
	ShipStatus      string    `json:"shippingstatus"`
}

type OrderInput struct {
	TransactionNo   string    `json:"transactno"`
	TransactionDate time.Time `json:"transacdate"`
	ItemCode        string    `json:"code"`
	Qty             int64     `json:"qty"`
	BasePrice       int64     `json:"baseprice"`
	SalePrice       int64     `json:"saleprice"`
	ShippingCost    int64     `json:"shippingprice"`
	Discount        int64     `json:"discount"`
	Total           int64     `json:"total"`
	AccountID       string    `json:"accid"`
	PaymentStatus   string    `json:"paymentstatus"`
	ShipStatus      string    `json:"shippingstatus"`
}
