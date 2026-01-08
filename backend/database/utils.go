package database

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func GetAllProduct(db *gorm.DB) ([]Product, error) {

	var alldata []Product;
	result := db.Find(&alldata);
	return alldata, result.Error;

}

func GetProduct(db *gorm.DB, id uuid.UUID) (Product, error) {

	product := Product{};
	result := db.Where("id = ?", id).Find(&product);
	if result.Error != nil {
		return product, fmt.Errorf("%s", result.Error);
	}

	return product, nil;

}

func AddProduct(db *gorm.DB, product Product) (error) {
	
	result := db.Create(&product); if result.Error != nil {
		return fmt.Errorf("error trying to add the register: %w", result.Error);
	} else {
		return nil;
	}

}

func UpdateProduct(db *gorm.DB, id uuid.UUID, product Product) (error) {

	result := db.Model(&Product{}).Where("id = ?", id).Updates(product);
	if result.Error != nil {
		return fmt.Errorf("error trying to update the data: %s", result.Error);
	}

	return nil;
}

func GetImage(db *gorm.DB, id uuid.UUID) (Image, error) {

	image := Image{};
	result := db.Where("id_image = ?", id).Find(&image);
	if result.Error != nil {
		return image, fmt.Errorf("%s", result.Error);
	}

	return image, nil;
}

func AddImage(db *gorm.DB, image Image) (error) {

	result := db.Create(&image); if result.Error != nil {
		return fmt.Errorf("error trying to add the image %s", result.Error);
	} else {
		return nil;
	}
}