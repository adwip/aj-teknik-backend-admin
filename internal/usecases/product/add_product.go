package product

import (
	"errors"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entity"
	"github.com/adwip/aj-teknik-backend-admin/internal/shared/utils"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/responses"
)

func (p *productImpl) CreateProduct(req requests.AddProductRequest) (out responses.AddProductResponse, err error) {

	category, err := p.categoryRepo.GetCategoryById(req.Category)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if category.SecureID == "" {
		err = errors.New("Category not Found")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	brand, err := p.brandRepo.GetBrandById(req.Brand)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if brand.SecureID == "" {
		err = errors.New("Brand not Found")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	newProduct := entity.Products{
		Name:     req.Name,
		SecureID: utils.GenerateUUID(),
		// Price:       req.Price,
		Stock:       req.Stock,
		Code:        req.Code,
		BrandId:     brand.SecureID,
		CategoryId:  category.SecureID,
		Description: req.Description,
	}

	err = p.productRepo.CreateProduct(newProduct)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	out = responses.AddProductResponse{
		SecureID: newProduct.SecureID,
		Name:     newProduct.Name,
	}
	return out, nil
}
