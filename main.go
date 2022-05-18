package main

import (
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
	"github.com/lyve_cloud/models"
)

func main(){

	// arg1 is the db driver, arg2 is the db file location
	// db, err := sql.Open("sqlite3", "./sqlite.db")
	models.NewMigration()
}