package mysql

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
	"github.com/adwip/aj-teknik-backend-admin/internal/model/entity"
)

type brandRepo struct {
	db *sql.DB
}

func SetupBrandRepository(db *sql.DB) model.Brands {
	return &brandRepo{
		db: db,
	}
}

func (b *brandRepo) CreateBrand(req entity.Brands) (err error) {
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
