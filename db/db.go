package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func Init() {

	conn, err := sql.Open("sqlite3", "csv_reader.db")

	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}

	db = conn

	if _, err := db.Exec(createTable); err != nil {
		log.Fatalln("Table create error: " + err.Error())
	}

}

func GetDB() *sql.DB {
	return db
}
