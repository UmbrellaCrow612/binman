package main

import (
	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/fetch"
	"github.com/UmbrellaCrow612/binman/cli/yml"
)

func main() {
	options := args.Parse()
	config := yml.Parse(options)

	for _, bin := range config.Binaries {
		fetch.FetchAndStoreBinary(&bin, options)
	}

	// call extract to extract the folder/key/[..zip/tar] into /bin then unlink downloads
}
