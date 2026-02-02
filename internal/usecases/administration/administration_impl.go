package administration

import "github.com/adwip/aj-teknik-backend-admin/internal/model"

type adm struct {
	brandsRepo model.Brands
}

func SetupAdministrationUsecase(brandsRepo model.Brands) Administration {
	return &adm{
		brandsRepo: brandsRepo,
	}
}
