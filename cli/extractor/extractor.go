package extractor

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/UmbrellaCrow612/binman/cli/args"
)

var supportedArchiveFormats = []string{".zip", ".tar", ".gz"}

// Gets all folders which are archive
func GetAllArchiveFiles(basePath string) ([]string, error) {
	downloadPath := filepath.Join(basePath, "downloads")
	var foundPaths []string

	_, err := os.Stat(downloadPath)
	if err != nil {
		return nil, err
	}

	err = filepath.WalkDir(downloadPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		extension := filepath.Ext(d.Name())
		if extension == "" {
			return nil
		}

		if !slices.Contains(supportedArchiveFormats, extension) {
			return nil
		}

		fullPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		foundPaths = append(foundPaths, fullPath)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return foundPaths, nil
}

// Extract all downloads
func Extract(options *args.Options) error {
	archPaths, err := GetAllArchiveFiles(options.Path)
	if err != nil {
		return err
	}

	for _, path := range archPaths {
		ext := filepath.Ext(path)
		var extractError error

		switch ext {
		case ".zip":
			extractError = unZip(path)
		case ".tar":
			extractError = extractTar(path)
		case ".gz":
			extractError = extractTarGz(path)
		default:
			extractError = fmt.Errorf("Unsupported  format %s ", ext)
		}

		if extractError != nil {
			return extractError
		}
	}

	return nil
}
