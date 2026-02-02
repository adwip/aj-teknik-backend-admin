package model

import "github.com/adwip/aj-teknik-backend-admin/internal/model/entity"

type Brands interface {
	CreateBrand(req entity.Brands) (err error)
	GetBrandById(secureId string) (out entity.Brands, err error)
}
