package product

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger/report"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
)

type productImpl struct {
	log          report.Report
	productRepo  model.Products
	brandRepo    model.Brands
	categoryRepo model.Category
}

func SetupProductUsecase(log report.Report, productRepo model.Products, brandRepo model.Brands, categoryRepo model.Category) Product {
	return &productImpl{
		log:          log,
		productRepo:  productRepo,
		brandRepo:    brandRepo,
		categoryRepo: categoryRepo,
	}
}
