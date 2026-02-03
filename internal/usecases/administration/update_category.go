package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

func (a *adm) UpdateCategory(req request.UpdateCategoryRequest) (out response.UpdateCategoryResponse, err error) {
	categoryentities := entities.Category{
		Name:        req.Name,
		Description: req.Description,
		Parent:      req.Parent,
	}

	err = a.categoryRepo.UpdateCategory(req.ID, categoryentities)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	out.Message = "Category updated successfully"
	return out, nil
}
