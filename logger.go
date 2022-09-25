package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct {
	file string
}

var logger *Logger = NewLogger()

func NewLogger() *Logger {
	l := Logger{
		file: path.Join(dirs.LOGS, fmt.Sprintf("log-%s", time.Now().Format(time.RFC3339))),
	}
	return &l
}

func (logger *Logger) Info(args ...string) {
	text := fmt.Sprintf("INFO: "+args[0], args[1:])
	logger.write(text)
}

func (logger *Logger) err(args ...string) {
	text := fmt.Sprintf("ERROR: "+args[0], args[1:])
	logger.write(text)
}

func (logger *Logger) write(text string) {
	f, err := os.OpenFile(logger.file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + text + "\n"); err != nil {
		log.Println(err)
	}
}
