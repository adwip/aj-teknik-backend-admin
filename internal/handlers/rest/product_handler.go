package rest

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
	"github.com/labstack/echo/v5"
)

type ProductHandler struct {
	product product.Product
}

func SetupProductHandler(product product.Product) *ProductHandler {
	return &ProductHandler{
		product: product,
	}
}

func (r *ProductHandler) GetProducts(c *echo.Context) error {

	req, err := requests.NewGetProductsRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	out, pagination, err := r.product.GetProducts(req)
	if err != nil {
		return session.SetResultWithPagination(c, nil, pagination, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResultWithPagination(c, out, pagination, nil)

}

func (r *ProductHandler) CreateProduct(c *echo.Context) error {
	req, err := requests.NewAddProductRequest(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, stacktrace.MESSAGE_INVALID_INPUT))
	}

	out, err := r.product.CreateProduct(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR))
	}
	return session.SetResult(c, out, nil)
}
