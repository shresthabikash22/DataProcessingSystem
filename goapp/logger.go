package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	infoLogger  = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stderr, "", 0)
)

func logInfo(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	infoLogger.Printf("%s [INFO] %s", timestamp, message)
}

func logError(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	errorLogger.Printf("%s [ERROR] %s", timestamp, message)
}

func fmtError(message string) error {
	return fmt.Errorf(message)
}