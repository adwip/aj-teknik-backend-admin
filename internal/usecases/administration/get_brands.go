package administration

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration/response"
)

func (a *adm) GetBrands() (out response.GetBrandsResponse, err error) {
	brands, err := a.brandsRepo.GetAllBrands()
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	var brandItems []response.BrandItem
	for _, brand := range brands {
		brandItems = append(brandItems, response.BrandItem{
			ID:          brand.SecureID,
			Name:        brand.Name,
			Description: brand.Description,
		})
	}

	out.Brands = brandItems
	return out, nil
}
