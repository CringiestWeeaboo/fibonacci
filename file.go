package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/goodsign/monday"
)

type FileWriter interface {
	createHistory(resultHistory []int) string
	writeHistoryToFile(history string)
	readLinesFromFile(file *os.File) (int, error)
}

type fileWriter struct {
	fileName string
}

func NewFileWriter(fileName string) FileWriter {
	return fileWriter{fileName: fileName}
}

func (f fileWriter) createHistory(resultHistory []int) string {
	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
	history := strconv.Itoa(resultHistory[len(resultHistory)-1]) + ", " + date
	return history
}

func (f fileWriter) writeHistoryToFile(history string) {
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal("Невозможно создать файл:", err)
	}

	lineCount, err := f.readLinesFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(strconv.Itoa(lineCount) + ". " + history + "\n")
	if err != nil {
		log.Fatal("Невозможно записать строку в файл:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal("Ошибка сброса буфера:", err)
		return
	}
}

func (f fileWriter) readLinesFromFile(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 1
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
