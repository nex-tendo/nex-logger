package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	file   *os.File
	logger *log.Logger
	mu     sync.Mutex
}

func NewLogger(filePath string) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}

	return &Logger{
		file:   file,
		logger: log.New(file, "", 0),
	}, nil
}

func (l *Logger) Close() {
	l.file.Close()
}

func (l *Logger) logMessage(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
	fmt.Println(logMsg)
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println(logMsg)
}

func (l *Logger) Info(message string) {
	l.logMessage("INFO", message)
}

func (l *Logger) Error(message string) {
	l.logMessage("ERROR", message)
}

func (l *Logger) Warning(message string) {
	l.logMessage("WARNING", message)
}

func (l *Logger) Success(message string) {
	l.logMessage("SUCCESS", message)
}
