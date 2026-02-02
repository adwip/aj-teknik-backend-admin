package sql_lib

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	yellow = "\033[33m"
	reset  = "\033[0m"
)

type DB struct {
	*sqlx.DB
}

func SetupSQL(url string) (*DB, error) {
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (s *DB) Select(dest interface{}, query string, args ...interface{}) (err error) {
	visualizeQuery(query, args)
	err = s.DB.Select(dest, query, args...)
	return err
}

func (s *DB) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	visualizeQuery(query, args)
	result, err = s.DB.Exec(query, args...)
	return result, err
}

func (s *DB) Get(dest interface{}, query string, args ...interface{}) (err error) {
	visualizeQuery(query, args)
	err = s.DB.Get(dest, query, args...)
	return err
}
