package printer

import (
	"os"

	"github.com/fatih/color"
)

// PrintSuccess prints a success message in green to stdout.
func PrintSuccess(msg string) {
	green := color.New(color.FgGreen)
	green.Fprintf(os.Stdout, "%s\n", msg)
}

// PrintError prints an error message in red to stderr.
func PrintError(msg string) {
	red := color.New(color.FgRed)
	red.Fprintf(os.Stderr, "%s\n", msg)
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
