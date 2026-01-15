package cleaner

import (
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/console"
	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Clean the project from folders and files downloaded and made by binman cli
// these include the downloads folder, extracted folder and the bin folder
func CleanStart(o *t.ArgOptions) error {
	paths := []string{
		filepath.Join(o.BasePath, "downloads"),
		filepath.Join(o.BasePath, "extracted"),
		filepath.Join(o.BasePath, "bin"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			console.LogInfo("Removing " + p)
		}

		if err := os.RemoveAll(p); err != nil {
			return err
		}
	}

	return nil
}

// CleanEnd removes only the downloads and extracted folders,
// leaving the bin folder intact
func CleanEnd(o *t.ArgOptions) error {
	paths := []string{
		filepath.Join(o.BasePath, "downloads"),
		filepath.Join(o.BasePath, "extracted"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			console.LogInfo("Removing " + p)
		}

		if err := os.RemoveAll(p); err != nil {
			return err
		}
	}

	return nil
}
