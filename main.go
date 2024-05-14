package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	var resultHistory []int
	mux := &http.ServeMux{}
	mux.HandleFunc("/fib/{num}", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, &resultHistory)
	})
	http.ListenAndServe(":8080", mux)

}

func indexHandler(w http.ResponseWriter, r *http.Request, resultHistory *[]int) {
	fibIter, err := strconv.Atoi(r.PathValue("num"))
	if errorMessage := validate(err, fibIter); errorMessage != nil {
		fmt.Fprint(w, errorMessage)
		return
	}
	fmt.Fprint(w, "Сумма: "+strconv.Itoa(fibonacci(fibIter))+"\n")
	fmt.Fprint(w, createHistoryMessage(*resultHistory))
	*resultHistory = append(*resultHistory, fibonacci(fibIter))
	writeHistoryToFile(createHistory(*resultHistory))
	queryToDatabase(fibIter, fibonacci(fibIter), createDatabaseAndTable())
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

func validate(err error, fibIterNum int) error {
	if err != nil {
		return errors.New("Ошибка! Введите число в адресной строке после 'fib/'")
	}
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
