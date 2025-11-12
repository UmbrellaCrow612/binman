package main

import (
	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/cleaner"
	"github.com/UmbrellaCrow612/binman/cli/fetch"
	"github.com/UmbrellaCrow612/binman/cli/yml"
)

func main() {
	options := args.Parse()
	config := yml.Parse(options)

	cleaner.Clean(options)

	for _, bin := range config.Binaries {
		fetch.FetchAndStoreBinary(&bin, options)
	}

	// cleaner.CleanDownloads(options)
}
