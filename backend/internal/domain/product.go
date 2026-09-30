package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID                 uuid.UUID
	ProductName        string
	ProductPrice       float32
	ProductDescription string
	IdImage            uuid.UUID
	CreatedAt          time.Time
	LastUpdate         time.Time
}
