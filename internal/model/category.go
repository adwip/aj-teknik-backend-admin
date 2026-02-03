package model

import "github.com/adwip/aj-teknik-backend-admin/internal/model/entities"

type Category interface {
	CreateCategory(req entities.Category) (err error)
	GetCategoryById(secureId string) (out entities.Category, err error)
	GetAllCategories() (out []entities.Category, err error)
	GetCategoryTree() (out []entities.Category, err error)
	UpdateCategory(secureId string, req entities.Category) (err error)
	DeleteCategory(secureId string) (err error)
}
