package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	IdUser           uuid.UUID
	UserName         string
	UserLogin        string
	UserPasswordHash string
	IdImage          string
	CreatedAt        time.Time
	LastUpdate       time.Time
}
