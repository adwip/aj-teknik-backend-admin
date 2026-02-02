package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

type Administration interface {
	AddBrand(req request.AddBrandRequest) (out response.AddBrandResponse, err error)
}
