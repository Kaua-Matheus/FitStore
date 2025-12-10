package database

import (
	"time"

	"github.com/google/uuid"
	_"gorm.io/gorm"
)

type Product struct {
	ID					uuid.UUID	`json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductName			string		`json:"product_name"`
	ProductPrice		float32		`json:"product_price"`
	ProductDescription	string		`json:"product_description"`
	IdImage				uuid.UUID	`json:"id_image" gorm:"type:uuid"`
	LastUpdate			time.Time	`json:"last_update" gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "product";
}

type ProductImage struct {
	IdImage		uuid.UUID	`json:"id_image" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ContentType	string		`json:"content_type"`
	FilePath	string		`json:"file_path"`
	FileName	string		`json:"file_name"`
}

func (ProductImage) TableName() string {
	return "product_image";
}