package extractor

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/console"
)

// Extract a .tar.gz file path to a dest file path
func ExtractTarGz(fromPath string, toPath string) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("from path is not a file: " + fromPath)
	}

	if filepath.Ext(fromPath) != ".gz" {
		return errors.New("from path file extension is not .gz: " + fromPath)
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

	if err := os.MkdirAll(toPath, 0755); err != nil {
		return err
	}

	destRoot := filepath.Clean(toPath) + string(os.PathSeparator)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(toPath, header.Name)
		targetPath = filepath.Clean(targetPath)

		if !strings.HasPrefix(targetPath, destRoot) {
			return errors.New("illegal file path: " + header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return err
			}

		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return err
			}

			outFile, err := os.OpenFile(
				targetPath,
				os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
				os.FileMode(header.Mode),
			)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()

		case tar.TypeSymlink:
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return err
			}
			if err := os.Symlink(header.Linkname, targetPath); err != nil {
				return err
			}

		default:
			console.LogWarning(fmt.Sprintf("Skipping unknown type: %c in %s\n", header.Typeflag, header.Name))
		}
	}

	return nil
}
