package config

import (
	"fmt"
	"log"

	"github.com/pintarkode/my_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", ENV.DATABASE_HOST, ENV.DATABASE_USER, ENV.DATABASE_PASSWORD, ENV.DATABASE_NAME, ENV.DATABASE_PORT, ENV.DATABASE_SSL)
	//dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disabled", ENV.DATABASE_HOST, ENV.DATABASE_USER, ENV.DATABASE_PASSWORD, ENV.DATABASE_NAME, ENV.DATABASE_PORT)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//call migration
	dbMigrate()

	//
	if err != nil {
		log.Fatal("Connection error...", err)
	}
	fmt.Println("Database Success Connnected")
}

func dbMigrate() {

	DB.AutoMigrate(
		models.Accounts{},
		models.Addresses{},
		models.Categories{},
		models.Items{},
		models.Accounts{},
		models.Orders{},
		models.Payments{},
		models.Brands{},
		models.ItemImage{},
		models.Promotions{},
		models.PromotionDetails{},
		models.Notification{},
		models.Company{},
		models.Message{},
	)
}
