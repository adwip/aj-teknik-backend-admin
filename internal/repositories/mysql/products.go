package mysql

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql_lib"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/dtos"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product/requests"
)

type productsRepo struct {
	db *sql_lib.DB
}

func SetupProductsRepository(db *sql_lib.DB) model.Products {
	return &productsRepo{
		db: db,
	}
}

func (p *productsRepo) GetProducts(in requests.GetProductsRequest) (out []dtos.ProductListResponseDto, pagination session.PaginationFormatter, err error) {

	query := "SELECT p.secure_id as secure_id, p.name as name, p.code as code, c.name as category, b.name as brand FROM products p INNER join categories c on p.category_id = c.secure_id INNER join brands b on p.brand_id = b.secure_id"
	args := []interface{}{}

	if in.Query != "" {
		query += " and p.name LIKE ?"
		args = append(args, "%"+in.Query+"%")
	}

	if in.Category != "" {
		query += " and p.category_id = ?"
		args = append(args, in.Category)
	}

	if in.Brand != "" {
		query += " and p.brand_id = ?"
		args = append(args, in.Brand)
	}

	if in.Availability {
		query += " and p.stock > 0"
	}

	if in.Sort != "" {
		query += " order by " + in.Sort
	}

	query += " limit ?"
	args = append(args, in.Limit)

	err = p.db.Select(&out, query, args...)
	if err != nil {
		return out, pagination, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, pagination, nil
}

func (p *productsRepo) CreateProduct(req entities.Products) (err error) {
	query := "INSERT INTO products (name, secure_id, price, stock, code, category_id, brand_id, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	args := []interface{}{
		req.Name,
		req.SecureID,
		// req.Price,
		req.Stock,
		req.Code,
		req.CategoryId,
		req.BrandId,
		req.Description,
	}

	_, err = p.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}

func (p *productsRepo) GetProductById(secureID string) (out entities.Products, err error) {
	query := "SELECT * FROM products WHERE secure_id = ?"
	err = p.db.Get(&out, query, secureID)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}
