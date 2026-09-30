package request

import (
	_ "time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	IdUser       uuid.UUID `gorm:"primaryKey"`
	UserName     string    `json:"user_name" gorm:"not null"`
	UserLogin    string    `json:"login" gorm:"not null"`
	UserPassword string    `json:"password" gorm:"not null"`
	// CreatedAt    time.Time `gorm:"autoUpdateTime"`
	// LastUpdate   time.Time `gorm:"autoUpdateTime"`
}

type LoginRequest struct {
	UserLogin    string `json:"login" gorm:"not null"`
	UserPassword string `json:"password" gorm:"not null"`
}
