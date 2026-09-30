package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/Kaua-Matheus/fitstore/backend/internal/config"
	repository "github.com/Kaua-Matheus/fitstore/backend/internal/repository"
)

// Only database connection

func NewConnection() (*gorm.DB, error) {

	// DSN
	dsn, err := config.InitializeEnv()

	// Database conn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("error opening connection: %s\n", err)
	}

	// Execute entities migrations
	err = db.AutoMigrate(
		&repository.Product{},
		&repository.User{},

		&repository.Image{}, // Remove future
	)
	if err != nil {
		fmt.Printf("error in automigrate: %s\n", err)
	}

	return db, nil
}
