package printer

import (
	"os"
	"time"

	"github.com/fatih/color"
)

// getTimestamp returns the current time formatted.
func getTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// PrintSuccess prints a success message in green to stdout with timestamp.
func PrintSuccess(msg string) {
	green := color.New(color.FgGreen)
	green.Fprintf(os.Stdout, "[%s] %s\n", getTimestamp(), msg)
}

// PrintError prints an error message in red to stderr with timestamp.
func PrintError(msg string) {
	red := color.New(color.FgRed)
	red.Fprintf(os.Stderr, "[%s] ERROR: %s\n", getTimestamp(), msg)
}

// PrintWarning prints a warning message in yellow to stdout with timestamp.
func PrintWarning(msg string) {
	yellow := color.New(color.FgYellow)
	yellow.Fprintf(os.Stdout, "[%s] WARNING: %s\n", getTimestamp(), msg)
}

// ExitSuccess prints a success message and exits with code 0.
func ExitSuccess(msg string) {
	PrintSuccess(msg)
	os.Exit(0)
}

// ExitError prints an error message and exits with code 1.
func ExitError(msg string) {
	PrintError(msg)
	os.Exit(1)
}
