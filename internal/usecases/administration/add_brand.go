package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entity"
	"github.com/adwip/aj-teknik-backend-admin/internal/shared/utils"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

func (a *adm) AddBrand(req request.AddBrandRequest) (out response.AddBrandResponse, err error) {

	newBrand := entity.Brands{
		Name:        req.Name,
		SecureID:    utils.GenerateUUID(),
		Description: req.Description,
	}

	err = a.brandsRepo.CreateBrand(newBrand)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return response.AddBrandResponse{
		SecureID: newBrand.SecureID,
		Name:     newBrand.Name,
	}, nil
}
