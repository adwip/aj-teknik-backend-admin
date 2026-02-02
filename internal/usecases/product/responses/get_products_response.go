package responses

type GetProductsResponse struct {
	Products []ProductResponse `json:"products"`
}

type ProductResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
