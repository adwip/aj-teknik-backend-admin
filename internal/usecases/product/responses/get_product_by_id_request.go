package responses

type GetProductByIdResponse struct {
	SecureID    string  `json:"secure_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	Code        string  `json:"code"`
	BrandId     string  `json:"brand_id"`
	CategoryId  string  `json:"category_id"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}
