package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

type Administration interface {
	AddBrand(req request.AddBrandRequest) (out response.AddBrandResponse, err error)
	AddCategory(req request.AddCategoryRequest) (out response.AddCategoryResponse, err error)

	// Brands
	GetBrands() (out response.GetBrandsResponse, err error)
	UpdateBrand(req request.UpdateBrandRequest) (out response.UpdateBrandResponse, err error)
	DeleteBrand(req request.DeleteBrandRequest) (err error)

	// Categories
	GetCategories() (out response.GetCategoriesResponse, err error)
	GetCategoryTree() (out response.GetCategoryTreeResponse, err error)
	UpdateCategory(req request.UpdateCategoryRequest) (out response.UpdateCategoryResponse, err error)
	DeleteCategory(req request.DeleteCategoryRequest) (err error)
}
