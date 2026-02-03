package administration

import (
	"errors"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
	"github.com/adwip/aj-teknik-backend-admin/internal/shared/utils"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

func (a *adm) AddCategory(req request.AddCategoryRequest) (out response.AddCategoryResponse, err error) {

	category, err := a.categoryRepo.GetCategoryById(req.Parent)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT)
	}

	if category.SecureID == "" {
		err = errors.New("Parent category not Found")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	newCategory := entities.Category{
		Name:        req.Name,
		SecureID:    utils.GenerateUUID(),
		Parent:      category.SecureID,
		Description: req.Description,
	}

	err = a.categoryRepo.CreateCategory(newCategory)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return response.AddCategoryResponse{
		SecureID: newCategory.SecureID,
		Name:     newCategory.Name,
	}, nil
}
