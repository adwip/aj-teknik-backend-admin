package rest

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/request"
	"github.com/labstack/echo/v5"
)

type AdmHandler struct {
	brand administration.Administration
}

func SetupAdministrationHandler(brand administration.Administration) *AdmHandler {
	return &AdmHandler{
		brand: brand,
	}
}

func (a *AdmHandler) AddBrand(c *echo.Context) error {
	req, err := request.NewAddBrandRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	out, err := a.brand.AddBrand(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}
