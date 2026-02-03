package requests

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type GetProductsRequest struct {
	Query        string `query:"query"`
	Category     string `query:"category"`
	Brand        string `query:"brand"`
	Availability bool   `query:"availability"`
	Sort         string `query:"sort"`
	Page         int    `query:"page"`
	Limit        int    `query:"limit"`
}

func NewGetProductsRequest(in *echo.Context) (out GetProductsRequest, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT)
	}
	return out, nil
}
