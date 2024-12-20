package database

import (
	"log"
	"myproject/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Database migration completed successfully")
}
