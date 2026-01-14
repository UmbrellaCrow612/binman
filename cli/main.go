package main

import (
	"fmt"

	"github.com/UmbrellaCrow612/binman/cli/arguments"
	"github.com/UmbrellaCrow612/binman/cli/console"
)

// Main entry point
func main() {
	options, err := arguments.Parse()
	if err != nil {
		console.ExitError(err)
	}

	fmt.Println(options)
}
