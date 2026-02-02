package product

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entity"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/responses"
)

func (p *productImpl) CreateProduct(req requests.AddProductRequest) (out responses.AddProductResponse, err error) {

	newProduct := entity.Products{
		Name: req.Name,
		Code: req.Code,
		// CategoryId:  req.CategoryId,
		// BrandId:     req.BrandId,
		Description: req.Description,
	}

	err = p.productRepo.CreateProduct(newProduct)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}
