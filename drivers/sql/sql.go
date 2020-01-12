package sql

import (
	"black_gate/config"
	"database/sql"
)

func Connect() (*sql.DB, error){
	db, err := sql.Open("mysql", config.Database().ConnectionString())
	if err != nil{
		return  db, err
	}
	err = db.Ping()
	return  db, err
}
