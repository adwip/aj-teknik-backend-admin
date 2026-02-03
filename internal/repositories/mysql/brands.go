package mysql

import (
	"database/sql"
	"errors"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql_lib"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entities"
)

type brandRepo struct {
	db *sql_lib.DB
}

func SetupBrandRepository(db *sql_lib.DB) model.Brands {
	return &brandRepo{
		db: db,
	}
}

func (b *brandRepo) CreateBrand(req entities.Brands) (err error) {
	query := "INSERT INTO brands (name, secure_id, description) VALUES (?, ?, ?)"
	args := []interface{}{
		req.Name,
		req.SecureID,
		req.Description,
	}

	_, err = b.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}

func (b *brandRepo) GetBrandById(secureId string) (out entities.Brands, err error) {
	query := "SELECT * FROM brands WHERE secure_id = ?"
	args := []interface{}{
		secureId,
	}

	if err = b.db.Get(&out, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return out, nil
		}
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}

func (b *brandRepo) GetAllBrands() (out []entities.Brands, err error) {
	query := "SELECT * FROM brands ORDER BY name ASC"

	if err = b.db.Select(&out, query); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return out, nil
}

func (b *brandRepo) UpdateBrand(secureId string, req entities.Brands) (err error) {
	query := "UPDATE brands SET name = ?, description = ? WHERE secure_id = ?"
	args := []interface{}{
		req.Name,
		req.Description,
		secureId,
	}

	_, err = b.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}

func (b *brandRepo) DeleteBrand(secureId string) (err error) {
	query := "DELETE FROM brands WHERE secure_id = ?"
	args := []interface{}{
		secureId,
	}

	_, err = b.db.Exec(query, args...)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, stacktrace.MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}
