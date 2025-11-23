package cleaner

import (
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
)

// CleanStart removes the "bin" and "downloads" folders at the beginning of the process.
func CleanStart(options *args.Options) error {
	dirs := []string{
		filepath.Join(options.Path, "bin"),
		filepath.Join(options.Path, "downloads"),
	}

	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
		printer.PrintSuccess("Removed " + dir)
	}
	return nil
}

// CleanEnd removes only the "downloads" folder at the end of the process.
func CleanEnd(options *args.Options) error {
	downloadsDir := filepath.Join(options.Path, "downloads")
	return os.RemoveAll(downloadsDir)
}
