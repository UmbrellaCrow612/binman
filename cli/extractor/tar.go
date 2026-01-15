package extractor

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/console"
)

// Extract a tar file path to a a destinaion path
func ExtractTar(fromPath string, toPath string) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("from path is not a file: " + fromPath)
	}

	ext := filepath.Ext(fromPath)
	if ext == "" || ext == "." {
		return errors.New("from path doesn't have a file extension: " + fromPath)
	}

	if ext != ".tar" {
		return errors.New("from path is not a tar file: " + fromPath)
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	tr := tar.NewReader(file)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(toPath, hdr.Name)

		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(hdr.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return err
			}

			outFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		default:
			console.LogWarning(fmt.Sprintf("Skipping unknown type: %c in %s\n", hdr.Typeflag, hdr.Name))
		}
	}

	return nil
}
