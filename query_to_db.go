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
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fibonacci (id INTEGER PRIMARY KEY AUTOINCREMENT, iteration INTEGER, sum_of_fibonacci INTEGER, date_and_time TEXT)")
	if err != nil {
		panic(err)
	}
	return db
}

func queryToDatabase(fibIter int, fibonacci int, db *sql.DB) {
	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
	_, err := db.Exec("INSERT INTO fibonacci (iteration, sum_of_fibonacci, date_and_time) VALUES (?, ?, ?)", fibIter, fibonacci, date)
	if err != nil {
		panic(err)
	}
	//TODO: вынести в будущем в main() закрытие бд с помощью defer db.Close()
}
