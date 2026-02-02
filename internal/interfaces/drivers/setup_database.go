package drivers

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql"
)

func SetupDatabase(url string) (db *sql.DB, err error) {
	db, err = sql.SetupSQL(url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
