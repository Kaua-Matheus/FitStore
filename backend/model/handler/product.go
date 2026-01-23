package handler

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
)

// Product
func GetAllProduct(db *gorm.DB) ([]entitie.Product, error) {

	var alldata []entitie.Product;
	result := db.Find(&alldata);
	return alldata, result.Error;

}

func GetProduct(db *gorm.DB, id uuid.UUID) (entitie.Product, error) {

	product := entitie.Product{};
	result := db.Where("id = ?", id).Find(&product);
	if result.Error != nil {
		return product, fmt.Errorf("%s", result.Error);
	}

	return product, nil;

}

func AddProduct(db *gorm.DB, product entitie.Product) (error) {
	
	result := db.Create(&product); if result.Error != nil {
		return fmt.Errorf("error trying to add the register: %w", result.Error);
	} else {
		return nil;
	}

}

func UpdateProduct(db *gorm.DB, id uuid.UUID, product entitie.Product) (error) {

	result := db.Model(&entitie.Product{}).Where("id = ?", id).Updates(product);
	if result.Error != nil {
		return fmt.Errorf("error trying to update the data: %s", result.Error);
	}

	return nil;
}