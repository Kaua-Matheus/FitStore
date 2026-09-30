package request

type CreateProductRequest struct {
	ProductName        string  `json:"product_name" binding:"required"`
	ProductPrice       float32 `json:"product_price" binding:"required,gt=0"`
	ProductDescription string  `json:"product_description"`
}
