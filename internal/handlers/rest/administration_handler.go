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

func (a *AdmHandler) AddCategory(c *echo.Context) error {
	req, err := request.NewAddCategoryRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	out, err := a.brand.AddCategory(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}

func (a *AdmHandler) GetBrands(c *echo.Context) error {
	out, err := a.brand.GetBrands()
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}

func (a *AdmHandler) UpdateBrand(c *echo.Context) error {
	req, err := request.NewUpdateBrandRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	out, err := a.brand.UpdateBrand(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}

func (a *AdmHandler) DeleteBrand(c *echo.Context) error {
	req, err := request.NewDeleteBrandRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	err = a.brand.DeleteBrand(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, nil, nil)
}

func (a *AdmHandler) GetCategories(c *echo.Context) error {
	out, err := a.brand.GetCategories()
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}

func (a *AdmHandler) GetCategoryTree(c *echo.Context) error {
	out, err := a.brand.GetCategoryTree()
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}

func (a *AdmHandler) UpdateCategory(c *echo.Context) error {
	req, err := request.NewUpdateCategoryRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	out, err := a.brand.UpdateCategory(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}

func (a *AdmHandler) DeleteCategory(c *echo.Context) error {
	req, err := request.NewDeleteCategoryRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	err = a.brand.DeleteCategory(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, nil, nil)
}
