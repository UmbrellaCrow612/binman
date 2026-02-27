package global

import (
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/t"
)

// Contains global paths used throughout the app

// Base path used to download content to default to empty
var BinmanDownloadPath string // global

// Base path for where extracted content will go
var BinmanExtractedPath string // global

// Base path for where the bin folder will be containg the binary's extracted
var BinmanBinPath string // global

// Creates base path used throughout the app globally using parsed arg config
func CreateGlobalBasePaths(argOptions *t.ArgOptions) {
	BinmanDownloadPath = filepath.Join(argOptions.BasePath, "binman_downloads")
	BinmanExtractedPath = filepath.Join(argOptions.BasePath, "binman_extracted")
	BinmanBinPath = filepath.Join(argOptions.BasePath, "bin")
}
