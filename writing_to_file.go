package main

import (
	"bufio"
	"fmt"
	"github.com/goodsign/monday"
	"log"
	"os"
	"strconv"
	"time"
)

func createHistory(resultHistory []int) string {
	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
	history := strconv.Itoa(resultHistory[len(resultHistory)-1]) + ", " + date
	return history
}

func writeHistoryToFile(history string) {
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal("Unable to create file:", err)
		os.Exit(1)
	}
	lineCount := readLinesFromFile(file)
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(strconv.Itoa(lineCount) + ". " + history + "\n")
	if err != nil {
		log.Fatal("Unable to write to file:", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal("Ошибка сброса буфера:", err)
		return
	}
}

func readLinesFromFile(file *os.File) int {
	scanner := bufio.NewScanner(file)
	lineCount := 1
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
		return 0 //Понимаю, что вряд ли так можно делать
	}
	return lineCount
}
