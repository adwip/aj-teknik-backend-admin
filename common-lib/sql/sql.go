package sql

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

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
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go visualizeQuery(&wg, query, args)
	err = s.DB.Select(dest, query, args...)
	return err
}

func (s *DB) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go visualizeQuery(&wg, query, args)
	result, err = s.DB.Exec(query, args...)
	return result, err
}

func visualizeQuery(wg *sync.WaitGroup, sql string, args []any) {
	result := sql
	for _, arg := range args {
		var v string
		switch a := arg.(type) {
		case string:
			v = "'" + a + "'"
		default:
			v = fmt.Sprint(a)
		}
		result = strings.Replace(result, "?", v, 1)
	}
	fmt.Println(yellow + result + reset)
	wg.Done()
}
