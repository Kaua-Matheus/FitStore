package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	domain "github.com/Kaua-Matheus/fitstore/backend/internal/domain"
)

// Product
func GetAllProduct(db *gorm.DB) ([]domain.Product, error) {

	var alldata []domain.Product
	result := db.Find(&alldata)
	return alldata, result.Error

}

func GetProduct(db *gorm.DB, id uuid.UUID) (domain.Product, error) {

	product := domain.Product{}
	result := db.Where("id = ?", id).Find(&product)
	if result.Error != nil {
		return product, fmt.Errorf("%s", result.Error)
	}

	return product, nil

}

func AddProduct(db *gorm.DB, product domain.Product) error {

	result := db.Create(&product)
	if result.Error != nil {
		return fmt.Errorf("error trying to add the register: %w", result.Error)
	} else {
		return nil
	}

}

func UpdateProduct(db *gorm.DB, id uuid.UUID, product domain.Product) error {

	result := db.Model(&domain.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return fmt.Errorf("error trying to update the data: %s", result.Error)
	}

	return nil
}
