package models

import (
	"log"

	"example.com/m/v2/config"
)

func RunMigrations() {
	err := config.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		log.Fatal("Failed to create uuid extension: %v", err)
	}

	err = config.DB.AutoMigrate(
		&Category{},
	)

	if err != nil{
		log.Fatal("Migration failed: %v", err)
	}

	log.Println("Migration successful")
}