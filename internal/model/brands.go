package model

import "github.com/adwip/aj-teknik-backend-admin/internal/model/entity"

type Brands interface {
	CreateBrand(req entity.Brands) (err error)
}
