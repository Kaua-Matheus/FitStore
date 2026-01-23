package handler

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
)

// Image
func GetImage(db *gorm.DB, id uuid.UUID) (entitie.Image, error) {

	image := entitie.Image{};
	result := db.Where("id_image = ?", id).Find(&image);
	if result.Error != nil {
		return image, fmt.Errorf("%s", result.Error);
	}

	return image, nil;
}

func AddImage(db *gorm.DB, image entitie.Image) (error) {

	result := db.Create(&image); if result.Error != nil {
		return fmt.Errorf("error trying to add the image %s", result.Error);
	} else {
		return nil;
	}
}