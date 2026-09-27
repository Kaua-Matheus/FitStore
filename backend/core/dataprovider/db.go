package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/core/domain"
)

func NewConnection() (*gorm.DB, error) {

	err := godotenv.Load(); if err != nil {
		fmt.Printf("error loading env: %s\n", err);
		fmt.Printf("using docker .env\n");
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),	
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
	);

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); if err != nil {
		fmt.Printf("error opening connection: %s\n", err);
	};

	// Faz a atualização das entidades no banco
	err = db.AutoMigrate(
		&models.Product{},
		&models.Image{},
		&models.User{},
	); if err != nil {
		fmt.Printf("error in automigrate: %s\n", err);
	};

	return db, nil;
}
