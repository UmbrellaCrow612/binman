package extractor

import (
	"archive/tar"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Extract a tar file path to a a destinaion path
func ExtractTar(fromPath string, toPath string) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("from path is not a file " + fromPath)
	}

	ext := filepath.Ext(fromPath)
	if ext == "" || ext == "." {
		return errors.New("from path doesnt have a file extension " + fromPath)
	}

	if ext != ".tar" {
		return errors.New("from path is not a tar file " + fromPath)
	}

	reader, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	tr := tar.NewReader(reader)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
