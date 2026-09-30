package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	repository "github.com/Kaua-Matheus/fitstore/backend/internal/repository"
)

// Image
func GetImage(db *gorm.DB, id uuid.UUID) (repository.Image, error) {

	image := repository.Image{}
	result := db.Where("id_image = ?", id).Find(&image)
	if result.Error != nil {
		return image, fmt.Errorf("%s", result.Error)
	}

	return image, nil
}

func AddImage(db *gorm.DB, image repository.Image) error {

	result := db.Create(&image)
	if result.Error != nil {
		return fmt.Errorf("error trying to add the image %s", result.Error)
	} else {
		return nil
	}
}
