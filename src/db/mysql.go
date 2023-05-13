package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	connect, err := sql.Open("mysql", "root:db1234@/recordings")
	if err != nil {
		return nil, err
	}
	if err = connect.Ping(); err != nil {
		return nil, err
	}

	return connect, nil
}
