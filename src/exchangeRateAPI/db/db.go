package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("subscribers.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB.AutoMigrate(&Subscriber{})
}

type Subscriber struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
}
