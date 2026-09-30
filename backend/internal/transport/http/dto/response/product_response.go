package response

type ProductResponse struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	Image        string  `json:"image_url,omitempty"`
}
