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
	LastUpdate			time.Time	`json:"last_update" gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "product";
}