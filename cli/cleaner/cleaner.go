package cleaner

import (
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
)

// clean deletes the bin and downloads folders inside the provided path
func Clean(opts *args.Options) {
	folders := []string{
		filepath.Join(opts.Path, "bin"),
		filepath.Join(opts.Path, "downloads"),
	}

	for _, folder := range folders {
		err := os.RemoveAll(folder)
		if err != nil {
			printer.PrintError("Failed to remove folder: " + folder + " - " + err.Error())
			continue
		}
		printer.PrintSuccess("Successfully removed folder: " + folder)
	}
}

// CleanDownloads deletes the downloads folder inside the provided path
func CleanDownloads(opts *args.Options) {
	downloadsPath := filepath.Join(opts.Path, "downloads")

	err := os.RemoveAll(downloadsPath)
	if err != nil {
		printer.PrintError("Failed to remove downloads folder: " + downloadsPath + " - " + err.Error())
		return
	}

	printer.PrintSuccess("Successfully removed downloads folder: " + downloadsPath)
}
