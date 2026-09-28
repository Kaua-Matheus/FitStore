package repository

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	IdUser           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserName         string    `gorm:"not null"`
	UserLogin        string    `gorm:"not null"`
	UserPasswordHash string    `gorm:"column:password_hash;not null"`
	IdImage          string    `gorm:"type:uuid"`
	CreatedAt        time.Time `gorm:"autoUpdateTime"`
	LastUpdate       time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}
