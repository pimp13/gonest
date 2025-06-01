package database

import (
	"github.com/pimp13/gonest/config"
	"github.com/pimp13/gonest/modules/users/models"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := config.NewDatabase().Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
}
