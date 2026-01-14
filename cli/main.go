package main

import (
	"github.com/UmbrellaCrow612/binman/cli/arguments"
	"github.com/UmbrellaCrow612/binman/cli/config"
	"github.com/UmbrellaCrow612/binman/cli/console"
)

// Main entry point
func main() {
	options, err := arguments.Parse()
	if err != nil {
		console.ExitError(err)
	}

	config, err := config.Parse()
}
