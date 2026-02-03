package response

type GetCategoriesResponse struct {
	Categories []CategoryItem `json:"categories"`
}

type CategoryItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}
