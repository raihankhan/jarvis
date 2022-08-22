package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CreateTable() {
	createSqlTable := `CREATE TABLE IF NOT EXISTS todo (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "task" TEXT
    );`

	statement, err := db.Prepare(createSqlTable)

	if err != nil {
		log.Println(err)
	}

	statement.Exec()
	log.Println("Successfully created todo table")
}
