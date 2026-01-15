package extractor

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Extract the the content of a zip file from to a destination path
// if the destination path does not exit it will create it
func extractZip(fromPath string, toPath string) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("path is a directory, not a zip file: " + fromPath)
	}

	if filepath.Ext(fromPath) != ".zip" {
		return errors.New("file is not a zip archive: " + fromPath)
	}

	if err := os.MkdirAll(toPath, os.ModePerm); err != nil {
		return err
	}

	archive, err := zip.OpenReader(fromPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	destRoot := filepath.Clean(toPath) + string(os.PathSeparator)

	for _, f := range archive.File {
		filePath := filepath.Join(toPath, f.Name)

		if !strings.HasPrefix(filePath, destRoot) {
			return errors.New("invalid file path: " + filePath)
		}

		if f.Mode()&os.ModeSymlink != 0 {
			return errors.New("symlinks are not allowed: " + f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		srcFile, err := f.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}
	}

	return nil
}
