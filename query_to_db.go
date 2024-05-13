package main

import (
	"database/sql"
	"github.com/goodsign/monday"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func createDatabaseAndTable() *sql.DB {
	db, err := sql.Open("sqlite3", "fibonacci.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fibonacci (" +
		"id INTEGER PRIMARY KEY, " +
		"iteration INTEGER, " +
		"sum_of_fibonacci INTEGER, " +
		"date_and_time TEXT)")
	if err != nil {
		panic(err)
	}
	return db
}

func queryToDatabase(fibIter int, fibonacci int, db *sql.DB) {
	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
	stmt, err := db.Prepare("INSERT INTO fibonacci VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(fibIter, fibonacci, date)
	if err != nil {
		panic(err)
	}
	//вынести в будущем в main() закрытие дб defer db.Close()
}
