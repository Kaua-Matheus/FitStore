package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
)

func NewConnection() (*gorm.DB, error) {

	err := godotenv.Load(); if err != nil {
		return nil, fmt.Errorf("error loading env: %s", err);
	}

	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    //     os.Getenv("DB_HOST"),
    //     os.Getenv("DB_USER"),
    //     os.Getenv("DB_PASSWORD"),
    //     os.Getenv("DB_NAME"),
    //     os.Getenv("DB_PORT"),
    //     os.Getenv("DB_SSLMODE"),
	// );

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); if err != nil {
		return nil, fmt.Errorf("error opening connection: %s", err);
	};

	// Faz a atualização das entidades no banco
	err = db.AutoMigrate(
		&entitie.Product{},
		&entitie.Image{},
		&entitie.User{},
	); if err != nil {
		return nil, fmt.Errorf("error in automigrate: %s", err);
	};

	return db, nil;
}
