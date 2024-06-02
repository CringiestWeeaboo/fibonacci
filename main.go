package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type handler struct {
	fibHistory []int
	db         *sql.DB
	mux        http.ServeMux
	fileWriter FileWriter
}

func (h *handler) handle(w http.ResponseWriter, r *http.Request, f fileWriter, d historyDatabase) {
	i, err := strconv.Atoi(r.PathValue("num"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err := validate(i); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "Сумма: "+strconv.Itoa(fibonacci(i))+"\n")
	fmt.Fprint(w, createHistoryMessage(h.fibHistory))
	h.fibHistory = append(h.fibHistory, fibonacci(i))
	f.writeHistoryToFile(f.createHistory(h.fibHistory))
	d.writeToDatabase(i, fibonacci(i))
}

func main() {
	var (
		h   = &handler{}
		f   fileWriter
		d   historyDatabase
		err error
	)
	h.fileWriter = NewFileWriter("history.txt")
	h.db, err = initDb("fibonacci.db")
	if err != nil {
		log.Fatalf("Невозможно открыть базу данных: %v", err)
	}
	defer h.db.Close()

	err = h.createTable()
	if err != nil {
		log.Fatalf("Невозможно создать таблицу: %v", err)
	}
	d.db = h.db
	h.mux.HandleFunc("/fib/{num}", func(w http.ResponseWriter, r *http.Request) {
		h.handle(w, r, f, d)
	})
	http.ListenAndServe(":8080", &h.mux)
}

func createHistoryMessage(resultHistory []int) string {
	var historyMessage string
	if len(resultHistory) > 0 {
		historyMessage += "\nИстория:"
	}
	for i := len(resultHistory) - 1; i >= 0; i-- {
		historyMessage += "\n" + strconv.Itoa(len(resultHistory)-i) + ". " + strconv.Itoa(resultHistory[i])
	}
	return historyMessage
}

func validate(fibIterNum int) error {
	if fibIterNum >= 45 || fibIterNum <= 0 {
		return errors.New("Ошибка! Введите число в адресной строке после 'fib/' меньше 45 и больше 0")
	}
	return nil
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
