package config

import (
	"GoREST/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitializeDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to config: ", err)
	}

	err = DB.AutoMigrate(&models.Album{}, &models.Artist{})
	if err != nil {
		log.Fatal("failed to migrate db")
	}
}
