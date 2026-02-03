package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

func (a *adm) UpdateBrand(req request.UpdateBrandRequest) (out response.UpdateBrandResponse, err error) {
	brandentities := entities.Brands{
		Name:        req.Name,
		Description: req.Description,
	}

	err = a.brandsRepo.UpdateBrand(req.ID, brandentities)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	out.Message = "Brand updated successfully"
	return out, nil
}
