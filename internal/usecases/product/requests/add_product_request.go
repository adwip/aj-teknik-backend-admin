package requests

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type AddProductRequest struct {
	Name           string           `json:"name" validate:"required"`
	Code           string           `json:"code" validate:"required"`
	Price          float64          `json:"price" validate:"required"`
	Stock          float64          `json:"stock" validate:"required"`
	Category       string           `json:"category" validate:"required"`
	Brand          string           `json:"brand" validate:"required"`
	Description    string           `json:"description" validate:"required"`
	Spesifications []Spesifications `json:"spesifications"`
	Images         string           `json:"images"`
}

type Spesifications struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewAddProductRequest(in *echo.Context) (out AddProductRequest, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT)
	}
	return out, nil
}
