package model

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/dtos"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
)

type Products interface {
	GetProducts(in requests.GetProductsRequest) (out []dtos.ProductListResponseDto, pagination session.PaginationFormatter, err error)
	CreateProduct(req entities.Products) (err error)
	GetProductById(secureID string) (out entities.Products, err error)
}
