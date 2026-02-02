package model

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entity"
)

type Products interface {
	GetProducts(name, category string, availability bool, sort string, page, limit int) (out []entity.Products, pagination session.PaginationFormatter, err error)
	CreateProduct(req entity.Products) (err error)
}
