package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_username     = "root"
	db_password     = "rootroot"
	db_address_port = "127.0.0.1:3306"
	db_database     = "jobportal"
)

func getDb() (*sql.DB, error) {
	myDataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", db_username, db_password, db_address_port, db_database)
	Db, err := sql.Open("mysql", myDataSource)
	if err != nil {
		return nil, err
	}
	return Db, nil
}
