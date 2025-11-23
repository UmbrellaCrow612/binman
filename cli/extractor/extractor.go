package extractor

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
)

// Will go to options.PATH/downloads and extract all the content to bin
// Assumes downloads looks like
// path/downloads/ripgrep/x86_64/ripgrep.zip
func Extract(options *args.Options) {
	downloads := filepath.Join(options.Path, "downloads")
	bin := filepath.Join(options.Path, "bin")

	// Ensure downloads exists
	if _, err := os.Stat(downloads); err != nil {
		printer.ExitError("downloads folder does not exist: " + err.Error())
	}

	os.RemoveAll(bin)
	if err := os.MkdirAll(bin, 0755); err != nil {
		printer.ExitError("failed to create bin directory: " + err.Error())
	}

	var archives []string

	// Collect all archives in downloads
	err := filepath.Walk(downloads, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		lower := strings.ToLower(info.Name())

		if strings.HasSuffix(lower, ".zip") ||
			strings.HasSuffix(lower, ".tar.gz") ||
			strings.HasSuffix(lower, ".tgz") ||
			strings.HasSuffix(lower, ".tar") {

			archives = append(archives, path)
		}

		return nil
	})

	if err != nil {
		printer.ExitError(err.Error())
	}

	// Extract each archive into matching bin path
	for _, archive := range archives {
		// Compute relative path inside downloads
		rel, err := filepath.Rel(downloads, archive)
		if err != nil {
			printer.ExitError("failed to compute relative path: " + err.Error())
		}

		// Folder inside bin (excluding file name)
		destDir := filepath.Dir(filepath.Join(bin, rel))
		if err := os.MkdirAll(destDir, 0755); err != nil {
			printer.ExitError("failed to create bin subdir: " + err.Error())
		}

		lower := strings.ToLower(archive)

		switch {
		case strings.HasSuffix(lower, ".zip"):
			err = extractZipAndDeleteTo(archive, destDir)

		case strings.HasSuffix(lower, ".tar.gz"), strings.HasSuffix(lower, ".tgz"):
			err = extractTarGzAndDeleteTo(archive, destDir)

		case strings.HasSuffix(lower, ".tar"):
			err = extractTarAndDeleteTo(archive, destDir)
		}

		if err != nil {
			printer.ExitError(err.Error())
		}
	}
}

func extractZipAndDeleteTo(zipPath, dest string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		extractPath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(extractPath, 0755); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(extractPath), 0755); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		out, err := os.Create(extractPath)
		if err != nil {
			rc.Close()
			return err
		}

		_, err = io.Copy(out, rc)
		rc.Close()
		out.Close()

		if err != nil {
			return err
		}
	}

	return os.Remove(zipPath)
}

func extractTarGzAndDeleteTo(tarGzPath, dest string) error {
	file, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	return extractTarTo(dest, gzReader, tarGzPath)
}

func extractTarAndDeleteTo(tarPath, dest string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return extractTarTo(dest, file, tarPath)
}

func extractTarTo(dest string, reader io.Reader, deletePath string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}

		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}

			out, err := os.Create(target)
			if err != nil {
				return err
			}

			_, err = io.Copy(out, tarReader)
			out.Close()

			if err != nil {
				return err
			}
		}
	}

	return os.Remove(deletePath)
}
