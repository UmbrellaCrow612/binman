package console

import (
	"fmt"
	"os"
	"time"
)

// ANSI color codes
const (
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
	reset  = "\033[0m" 
)

// logMessage prints a structured message with timestamp and color
func logMessage(level string, color string, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// Color only the level, then reset
	fmt.Printf("[%s] %s%s%s %s\n", timestamp, color, level, reset, msg)
}

// LogInfo prints an info message in green
func LogInfo(msg string) {
	logMessage("INFO", green, msg)
}

// LogWarning prints a warning message in yellow
func LogWarning(msg string) {
	logMessage("WARN", yellow, msg)
}

// LogError prints an error message in red
func LogError(msg string) {
	logMessage("ERROR", red, msg)
}

// ExitError prints the error message and exits with status 1
func ExitError(err error) {
	if err != nil {
		LogError(err.Error())
	}
	os.Exit(1)
}