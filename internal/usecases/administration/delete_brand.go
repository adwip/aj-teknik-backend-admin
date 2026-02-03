package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
)

func (a *adm) DeleteBrand(req request.DeleteBrandRequest) (err error) {
	err = a.brandsRepo.DeleteBrand(req.ID)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}
