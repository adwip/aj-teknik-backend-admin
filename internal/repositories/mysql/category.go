package mysql

import (
	"database/sql"
	"errors"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql_lib"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
)

type categoryRepo struct {
	db *sql_lib.DB
}

func SetupCategoryRepository(db *sql_lib.DB) model.Category {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) CreateCategory(req entities.Category) (err error) {
	query := "INSERT INTO categories (name, secure_id, description) VALUES (?, ?, ?)"
	args := []interface{}{
		req.Name,
		req.SecureID,
		req.Description,
	}

	_, err = c.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}

func (c *categoryRepo) GetCategoryById(secureId string) (out entities.Category, err error) {
	query := "SELECT * FROM categories WHERE secure_id = ?"
	args := []interface{}{
		secureId,
	}

	if err = c.db.Get(&out, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return out, nil
		}
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}

func (c *categoryRepo) GetAllCategories() (out []entities.Category, err error) {
	query := "SELECT * FROM categories ORDER BY name ASC"

	if err = c.db.Select(&out, query); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}

func (c *categoryRepo) GetCategoryTree() (out []entities.Category, err error) {
	query := "SELECT * FROM categories ORDER BY parent ASC, name ASC"

	if err = c.db.Select(&out, query); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}

func (c *categoryRepo) UpdateCategory(secureId string, req entities.Category) (err error) {
	query := "UPDATE categories SET name = ?, description = ?, parent = ? WHERE secure_id = ?"
	args := []interface{}{
		req.Name,
		req.Description,
		req.Parent,
		secureId,
	}

	_, err = c.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}

func (c *categoryRepo) DeleteCategory(secureId string) (err error) {
	query := "DELETE FROM categories WHERE secure_id = ?"
	args := []interface{}{
		secureId,
	}

	_, err = c.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}
