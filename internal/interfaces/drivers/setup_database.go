package drivers

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/sql_lib"
)

func SetupDatabase(url string) (db *sql_lib.DB, err error) {
	db, err = sql_lib.SetupSQL(url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
