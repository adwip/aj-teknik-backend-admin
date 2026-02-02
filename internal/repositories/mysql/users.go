package mysql

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql_lib"
	"github.com/adwip/aj-teknik-backend-admin/internal/model"
)

type usersRepository struct {
	db *sql_lib.DB
}

func SetupUsersRepository(db *sql_lib.DB) model.Users {
	return &usersRepository{
		db: db,
	}
}
