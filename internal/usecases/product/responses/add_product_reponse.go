package responses

type AddProductResponse struct {
	SecureID string `json:"secure_id"`
	Name     string `json:"name"`
}
