package extractor

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"os"
	"path/filepath"
)

// Extract a .tar.gz file path to a dest file path
func ExtractTarGz(fromPath string, toPath string) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("from path is not a file " + fromPath)
	}

	ext := filepath.Ext(fromPath)
	if ext == "" || ext == "." {
		return errors.New("from path does not have a file extension " + fromPath)
	}

	if ext != ".gz" {
		return errors.New("from path file extension is not a .gz extension " + fromPath)
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	return nil
}
