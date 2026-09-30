package request

import "github.com/google/uuid"

type CreateProductRequest struct {
	ID                 uuid.UUID `gorm:"primaryKey"`
	ProductName        string    `json:"product_name" binding:"required"`
	ProductPrice       float32   `json:"product_price" binding:"required,gt=0"`
	ProductDescription string    `json:"product_description"`
	IdImage            uuid.UUID `json:"id_image"`
}
