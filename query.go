package main

import (
	"database/sql"
	"fmt"
	"github.com/goodsign/monday"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type DbWriter interface {
	writeToDatabase(i int, fibonacci int)
}

type historyDatabase struct {
	db *sql.DB
}

func (database historyDatabase) writeToDatabase(i int, fibonacci int) {
	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
	_, err := database.db.Exec("INSERT INTO fibonacci (iteration, sum_of_fibonacci, date_and_time) VALUES (?, ?, ?)", i, fibonacci, date)
	if err != nil {
		log.Fatalf("Невозможно создать таблицу: %v", err)
	}
}

func NewDbWriter(db *sql.DB) DbWriter {
	return historyDatabase{db: db}
}

func initDb(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Невозможно создать базу данных: %w", err)
	}

	return db, nil
}

func (h handler) createTable() error {
	_, err := h.db.Exec("CREATE TABLE IF NOT EXISTS fibonacci (id INTEGER PRIMARY KEY AUTOINCREMENT, iteration INTEGER, sum_of_fibonacci INTEGER, date_and_time TEXT)")
	if err != nil {
		return fmt.Errorf("Невозможно создать базу данных: %w", err)
	}

	return nil
}

//func writeToDatabase(i int, fibonacci int, db *sql.DB) {
//	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
//	_, err := db.Exec("INSERT INTO fibonacci (iteration, sum_of_fibonacci, date_and_time) VALUES (?, ?, ?)", i, fibonacci, date)
//	if err != nil {
//		log.Fatalf("Невозможно создать таблицу: %v", err)
//	}
//}
