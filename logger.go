package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Info(message string) {
	logMessage("INFO", message)
}

func Error(message string) {
	logMessage("ERROR", message)
}

func Warning(message string) {
	logMessage("WARNING", message)
}

func Success(message string) {
	logMessage("SUCCESS", message)
}

func logMessage(level, message string) {
	fmt.Printf("[%s] %s: %s\n", getCurrentTime(), level, message)
}

func LogToFile(level, message string) {
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", 0)
	logger.Printf("[%s] %s: %s\n", getCurrentTime(), level, message)
}
