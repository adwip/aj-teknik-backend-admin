package request

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type AddCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}

func NewAddCategoryRequest(in *echo.Context) (out AddCategoryRequest, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT)
	}
	return out, nil
}
