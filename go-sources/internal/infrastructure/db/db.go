package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func open() *sql.DB {
	db, err := sql.Open("sqlite3", "../dblite/db.db")
	checkErr(err)
	defer func(db *sql.DB) {
		err := db.Close()
		checkErr(err)
	}(db)
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
