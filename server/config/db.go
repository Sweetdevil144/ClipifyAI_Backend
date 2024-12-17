package config

import (
	"ClipifyAI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := Config("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment file")
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to NeonDB: %v", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(&models.User{}, &models.Short{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	DB = db
	log.Println("Successfully connected to NeonDB")
}
