package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Taeil3714"
	dbName := "final_task_pbi"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return

}
