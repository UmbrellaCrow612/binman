package extractor

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Unzips a zip file into a folder in the same directory
func unZip(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return errors.New("path is a directory, not a zip file")
	}

	if !strings.HasSuffix(strings.ToLower(path), ".zip") {
		return errors.New("file is not a zip")
	}

	r, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer r.Close()

	destDir := strings.TrimSuffix(path, filepath.Ext(path))

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for _, f := range r.File {
		fPath := filepath.Join(destDir, f.Name)

		if !strings.HasPrefix(fPath, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return errors.New("illegal file path in zip")
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fPath, f.Mode()); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
