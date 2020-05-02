package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_template/config"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Database().ConnectionString())
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}
