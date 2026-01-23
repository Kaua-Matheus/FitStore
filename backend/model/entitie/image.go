package entitie

import (
	"github.com/google/uuid"
)

type Image struct {
	IdImage     uuid.UUID `json:"id_image" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ContentType string    `json:"content_type"`
	FilePath    string    `json:"file_path"`
	FileName    string    `json:"file_name"`
}

func (Image) TableName() string {
	return "image"
}