package request

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type UpdateCategoryRequest struct {
	ID          string `param:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}

func NewUpdateCategoryRequest(in *echo.Context) (out UpdateCategoryRequest, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT)
	}

	if err = in.Validate(out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT)
	}
	return out, nil
}
