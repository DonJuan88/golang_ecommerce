package models

import (
	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Uuid         string `json:"uid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	AccountEmail string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	IsAdmin      bool   `json:"admin"`
	Active       bool   `json:"active" gorm:"default:true"`
}

type Register struct {
	FirstName            string `json:"firstname"`
	LastName             string `json:"lastname"`
	Email                string `json:"email" `
	Password             string `json:"password" `
	PasswordConfirmation string `json:"password_confirmation" `
	Points               int    `json:"points"`
	IsAdmin              bool   `json:"admin"`
	Active               bool   `json:"active" gorm:"default:false"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthAccount struct {
	ID       int    `json:"ID"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"admin"`
	jwt.RegisteredClaims
}

type UpdateAccount struct {
	LastPassword         string `json:"last_password"`
	Password             string `json:"password" `
	PasswordConfirmation string `json:"password_confirmation" `
	Active               bool   `json:"active"`
}

type AccountActivation struct {
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

type EmailPasswordReset struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPassword struct {
        Token       string `json:"token" binding:"required"`
        NewPassword string `json:"new_password" binding:"required,min=8"`
    }

type UpdatePassword struct{
	Email string `json:"email" binding:"required,email"`
    NewPassword string `json:"new_password" binding:"required,min=8"`
	
}