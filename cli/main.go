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

	conf, err := config.Parse(options.ConfigPath)
	if err != nil {
		console.ExitError(err)
	}

	err = config.Validate(conf)
	if err != nil {
		console.ExitError(err)
	}
}
