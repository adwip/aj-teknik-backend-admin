package model

import "github.com/adwip/aj-teknik-backend-admin/internal/model/entity"

type Category interface {
	CreateCategory(req entity.Category) (err error)
	GetCategoryById(secureId string) (out entity.Category, err error)
	GetAllCategories() (out []entity.Category, err error)
	GetCategoryTree() (out []entity.Category, err error)
	UpdateCategory(secureId string, req entity.Category) (err error)
	DeleteCategory(secureId string) (err error)
}
