package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type handler struct {
	fibHistory []int
}

func (h *handler) indexHandler(w http.ResponseWriter, r *http.Request, fibHistory *[]int) {
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
	fmt.Fprint(w, createHistoryMessage(*fibHistory))
	*fibHistory = append(*fibHistory, fibonacci(i))
	writeHistoryToFile(createHistory(*fibHistory))
	queryToDatabase(i, fibonacci(i), createDatabaseAndTable())
}

func main() {
	mux := &http.ServeMux{}
	h := &handler{}
	mux.HandleFunc("/fib/{num}", func(w http.ResponseWriter, r *http.Request) {
		h.indexHandler(w, r, &h.fibHistory)
	})
	http.ListenAndServe(":8080", mux)

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
