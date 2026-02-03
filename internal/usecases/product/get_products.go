package product

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/responses"
)

func (p *productImpl) GetProducts(req requests.GetProductsRequest) (out responses.GetProductsResponse, pagination session.PaginationFormatter, err error) {

	products, pagination, err := p.productRepo.GetProducts(req)
	if err != nil {
		return out, pagination, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}

	for _, product := range products {
		out.Products = append(out.Products, responses.ProductResponse{
			SecureID: product.SecureID,
			Name:     product.Name,
			Image:    product.Image,
			Code:     product.Code,
			Category: product.Category,
			Brand:    product.Brand,
			// Price: product.Price,
			Stock: product.Stock,
		})
	}

	return out, pagination, nil
}
