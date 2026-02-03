package model

import "github.com/adwip/aj-teknik-backend-admin/internal/model/entities"

type Brands interface {
	CreateBrand(req entities.Brands) (err error)
	GetBrandById(secureId string) (out entities.Brands, err error)
	GetAllBrands() (out []entities.Brands, err error)
	UpdateBrand(secureId string, req entities.Brands) (err error)
	DeleteBrand(secureId string) (err error)
}
