package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID
	UserName         string
	UserLogin        string
	UserPasswordHash string
	IdImage          uuid.UUID
	CreatedAt        time.Time
	LastUpdate       time.Time
}
