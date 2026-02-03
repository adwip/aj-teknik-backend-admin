package response

type GetBrandsResponse struct {
	Brands []BrandItem `json:"brands"`
}

type BrandItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
