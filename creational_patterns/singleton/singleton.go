package main

import (
	"fmt"
	"sync"
)

// MyLogger is the struct we want to make a singleton.
type MyLogger struct {
	logLevel int
}

// Log a message using the Logger.
func (l *MyLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

// SetLogLevel sets the log level of the Logger.
func (l *MyLogger) SetLogLevel(level int) {
	l.logLevel = level
}

// the Logger instance.
var logger *MyLogger

// sync package use to enforce goroutine safety.
var once sync.Once

// getLoggerInstance provides global access to the logger type instance.
func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	})
	fmt.Println("Returning Logger Instance")

	return logger
}
