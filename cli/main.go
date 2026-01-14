package main

import (
	"fmt"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/arguments"
	"github.com/UmbrellaCrow612/binman/cli/config"
	"github.com/UmbrellaCrow612/binman/cli/console"
	"github.com/UmbrellaCrow612/binman/cli/global"
)

// Main entry point
func main() {
	options, err := arguments.Parse()
	if err != nil {
		console.ExitError(err)
	}

	if global.Verbose {
		console.LogInfo("Base path: " + options.BasePath)
		console.LogInfo("Binman yml file path: " + options.ConfigPath)
		console.LogInfo("Build platforms: " + strings.Join(options.Platforms, ","))
		console.LogInfo("Build architectures: " + strings.Join(options.Architectures, ","))

	}

	fmt.Println(options)

	conf, err := config.Parse(options.ConfigPath)
	if err != nil {
		console.ExitError(err)
	}

	err = config.Validate(conf)
	if err != nil {
		console.ExitError(err)
	}

	// for _, pack := range *conf {
	// 	fetch.Get(&pack, options)
	// }
}
