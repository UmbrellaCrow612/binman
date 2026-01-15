package extractor

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path"
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
		return errors.New("Path is not a valid zip path " + fromPath)
	}
	ext := path.Ext(fromPath)
	if ext == "" || ext != ".zip" {
		return errors.New("Path is not a zip path " + fromPath)
	}

	archive, err := zip.OpenReader(fromPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(toPath, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(toPath)+string(os.PathSeparator)) {
			return errors.New("invalid file path " + filePath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}
