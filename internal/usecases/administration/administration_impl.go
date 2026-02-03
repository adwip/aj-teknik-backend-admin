package administration

import "github.com/adwip/aj-teknik-backend-admin/internal/model"

type adm struct {
	brandsRepo   model.Brands
	categoryRepo model.Category
}

func SetupAdministrationUsecase(brandsRepo model.Brands, categoryRepo model.Category) Administration {
	return &adm{
		brandsRepo:   brandsRepo,
		categoryRepo: categoryRepo,
	}
}
