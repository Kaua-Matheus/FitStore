package mapper

import (
	domain "github.com/Kaua-Matheus/fitstore/backend/internal/domain"
	response "github.com/Kaua-Matheus/fitstore/backend/internal/transport/http/dto/response"
)

func ToProductResponse(p domain.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:           p.ID.String(),
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}
