package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

func (a *adm) GetCategories() (out response.GetCategoriesResponse, err error) {
	categories, err := a.categoryRepo.GetAllCategories()
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	var categoryItems []response.CategoryItem
	for _, category := range categories {
		categoryItems = append(categoryItems, response.CategoryItem{
			ID:          category.SecureID,
			Name:        category.Name,
			Description: category.Description,
			Parent:      category.Parent,
		})
	}

	out.Categories = categoryItems
	return out, nil
}
