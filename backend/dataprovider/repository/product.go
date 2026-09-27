package repository

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/core/domain"
)

// Product
func GetAllProduct(db *gorm.DB) ([]models.Product, error) {

	var alldata []models.Product;
	result := db.Find(&alldata);
	return alldata, result.Error;

}

func GetProduct(db *gorm.DB, id uuid.UUID) (models.Product, error) {

	product := models.Product{};
	result := db.Where("id = ?", id).Find(&product);
	if result.Error != nil {
		return product, fmt.Errorf("%s", result.Error);
	}

	return product, nil;

}

func AddProduct(db *gorm.DB, product models.Product) (error) {
	
	result := db.Create(&product); if result.Error != nil {
		return fmt.Errorf("error trying to add the register: %w", result.Error);
	} else {
		return nil;
	}

}

func UpdateProduct(db *gorm.DB, id uuid.UUID, product models.Product) (error) {

	result := db.Model(&models.Product{}).Where("id = ?", id).Updates(product);
	if result.Error != nil {
		return fmt.Errorf("error trying to update the data: %s", result.Error);
	}

	return nil;
}