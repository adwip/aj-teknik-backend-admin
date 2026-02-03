package product

import (
	"errors"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/responses"
)

func (s *productImpl) GetProductById(req requests.GetProductByIdRequest) (out responses.GetProductByIdResponse, err error) {
	product, err := s.productRepo.GetProductById(req.SecureID)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if product.SecureID == "" {
		err = errors.New("Product not Found")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	out = responses.GetProductByIdResponse{
		SecureID: product.SecureID,
		Name:     product.Name,
		// Price:       product.Price,
		Stock:       product.Stock,
		Code:        product.Code,
		BrandId:     product.BrandId,
		CategoryId:  product.CategoryId,
		Description: product.Description,
		Image:       product.Image,
	}
	return out, nil
}
