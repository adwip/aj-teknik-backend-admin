package product

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/responses"
)

type Product interface {
	GetProducts(req requests.GetProductsRequest) (out responses.GetProductsResponse, pagination session.PaginationFormatter, err error)
	CreateProduct(req requests.AddProductRequest) (out responses.AddProductResponse, err error)
}
