package main

import (
	"strings"
	"time"

	"github.com/UmbrellaCrow612/binman/arguments"
	"github.com/UmbrellaCrow612/binman/cleaner"
	"github.com/UmbrellaCrow612/binman/config"
	"github.com/UmbrellaCrow612/binman/console"
	"github.com/UmbrellaCrow612/binman/extractor"
	"github.com/UmbrellaCrow612/binman/fetch"
	"github.com/UmbrellaCrow612/binman/global"
	"github.com/UmbrellaCrow612/binman/transfer"
)

// Main entry point
func main() {
	start := time.Now()

	options, err := arguments.Parse()
	if err != nil {
		console.ExitError(err)
	}

	console.LogInfo("Base path: " + options.BasePath)
	console.LogInfo("Binman yml file path: " + options.ConfigPath)
	console.LogInfo("Build platforms: " + strings.Join(options.Platforms, ","))
	console.LogInfo("Build architectures: " + strings.Join(options.Architectures, ","))
	console.LogInfo("Build packages: " + strings.Join(options.Packages, ", "))

	global.CreateGlobalBasePaths(options)

	cleaner.CleanStart(options)

	conf, err := config.Parse(options.ConfigPath)
	if err != nil {
		console.ExitError(err)
	}

	err = config.Validate(conf)
	if err != nil {
		console.ExitError(err)
	}

	err = arguments.ValidateWithConfig(options, conf)
	if err != nil {
		console.ExitError(err)
	}

	for _, pack := range *conf {
		err := fetch.Get(&pack, options)
		if err != nil {
			console.ExitError(err)

		}
	}

	for _, pack := range *conf {
		err := extractor.Extract(&pack, options)
		if err != nil {
			console.ExitError(err)
		}
	}

	for _, pack := range *conf {
		err := transfer.ToBin(&pack, options)
		if err != nil {
			console.ExitError(err)
		}
	}

	cleaner.CleanEnd(options)

	elapsed := time.Since(start)
	console.LogInfo("Total time taken: " + elapsed.String())
}
