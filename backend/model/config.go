package model

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Product struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductName        string    `json:"product_name"`
	ProductPrice       float32   `json:"product_price"`
	ProductDescription string    `json:"product_description"`
	IdImage            uuid.UUID `json:"id_image" gorm:"type:uuid"`
	LastUpdate         time.Time `json:"last_update" gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "product"
}

type Image struct {
	IdImage     uuid.UUID `json:"id_image" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ContentType string    `json:"content_type"`
	FilePath    string    `json:"file_path"`
	FileName    string    `json:"file_name"`
}

func (Image) TableName() string {
	return "image"
}

type User struct {
	IdUser           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserName         string    `gorm:"not null"`
	UserLogin        string    `gorm:"not null"`
	UserPasswordHash string    `gorm:"column:password_hash;not null"`
	IdImage          string    `json:"id_image" gorm:"type:uuid"`
	CreatedAt        time.Time `gorm:"autoUpdateTime"`
	LastUpdate       time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}

type UserReq struct { // Struct somente utilizado para requisições json
	UserName     string `json:"user_name" gorm:"not null"`
	UserLogin    string `json:"login" gorm:"not null"`
	UserPassword string `json:"password" gorm:"not null"` // Adicionar criptografia na troca
}
