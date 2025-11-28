package main

import (
	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/cleaner"
	"github.com/UmbrellaCrow612/binman/cli/extractor"
	"github.com/UmbrellaCrow612/binman/cli/fetch"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/yml"
)

func main() {
	options := args.Parse()
	config := yml.Parse(options)

	cleaner.CleanStart(options)

	for _, bin := range config.Binaries {
		err := fetch.FetchAndStoreBinary(&bin, options)
		if err != nil {
			printer.ExitError(err.Error())

		}
	}

	err := extractor.Extract(options)
	if err != nil {
		printer.ExitError(err.Error())
	}

	// for _, bin := range config.Binaries {
	// 	err := extractor.CopyToBin(&bin, options)
	// 	if err != nil {
	// 		printer.ExitError(err.Error())
	// 	}
	// }

	// if !options.NoClean {
	// 	for _, bin := range config.Binaries {
	// 		err := pattern.CleanWithPattern(&bin, options)
	// 		if err != nil {
	// 			printer.ExitError(err.Error())
	// 		}
	// 	}
	// } else {
	// 	printer.PrintSuccess("No clean enabled skipping clean")
	// }

	// cleaner.CleanEnd(options)
}
