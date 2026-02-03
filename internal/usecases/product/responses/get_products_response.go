package responses

type GetProductsResponse struct {
	Products []ProductResponse `json:"products"`
}

type ProductResponse struct {
	SecureID    string  `json:"secure_id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
}
