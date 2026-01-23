package entitie

import (
	"time"
	"github.com/google/uuid"
)

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