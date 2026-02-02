package mysql

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
)

type usersRepository struct {
	db *sql.DB
}

func SetupUsersRepository(db *sql.DB) model.Users {
	return &usersRepository{
		db: db,
	}
}
