package repository

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductName        string    `gorm:"not null"`
	ProductPrice       float32   `gorm:"not null"`
	ProductDescription string    `gorm:"not_null"`
	IdImage            uuid.UUID `gorm:"type:uuid"`
	CreatedAt          time.Time `gorm:"autoUpdateTime"`
	LastUpdate         time.Time `gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "product"
}
