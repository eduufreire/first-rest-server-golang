package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:urubu100@/rest_server")
	if err != nil {
		fmt.Printf("unable to access database: %s", err)
	}
	return db
}
