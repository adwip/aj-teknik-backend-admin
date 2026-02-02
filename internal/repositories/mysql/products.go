package mysql

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql_lib"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entity"
)

type productsRepo struct {
	db *sql_lib.DB
}

func SetupProductsRepository(db *sql_lib.DB) model.Products {
	return &productsRepo{
		db: db,
	}
}

func (p *productsRepo) GetProducts(name, category string, availability bool, sort string, page, limit int) (out []entity.Products, pagination session.PaginationFormatter, err error) {

	query := "SELECT * FROM products where name LIKE ?"
	args := []interface{}{"%" + name + "%"}

	if category != "" {
		query += " and category_id = ?"
		args = append(args, category)
	}

	if availability {
		query += " and stock > 0"
	}

	if sort != "" {
		query += " order by " + sort
	}

	query += " limit ? offset ?"
	args = append(args, limit, (page-1)*limit)

	err = p.db.Select(&out, query, args...)
	if err != nil {
		return out, pagination, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, pagination, nil
}

func (p *productsRepo) CreateProduct(req entity.Products) (err error) {
	query := "INSERT INTO products (name, secure_id, price, stock, code, category_id, brand_id, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	args := []interface{}{
		req.Name,
		req.SecureID,
		req.Price,
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
