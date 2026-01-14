package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {

	err := godotenv.Load(); if err != nil {
		return nil, fmt.Errorf("error loading env: %s", err);
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
		return nil, fmt.Errorf("error opening connection: %s", err);
	};

	err = db.AutoMigrate(
		&Product{},
		&Image{},
		&User{},
	); if err != nil {
		return nil, fmt.Errorf("error in automigrate: %s", err);
	};

	return db, nil;
}
