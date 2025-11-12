package extractor

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
)

// Extract ZIP file to destination
func extractZip(src, dest string) error {
	printer.PrintSuccess(fmt.Sprintf("Extracting ZIP: %s -> %s", src, dest))
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fPath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.Create(fPath)
		if err != nil {
			inFile.Close()
			return err
		}

		if _, err := io.Copy(outFile, inFile); err != nil {
			inFile.Close()
			outFile.Close()
			return err
		}

		inFile.Close()
		outFile.Close()
	}

	return flattenSingleFolder(dest)
}

// Extract TAR.GZ file to destination
func extractTarGz(src, dest string) error {
	printer.PrintSuccess(fmt.Sprintf("Extracting TAR.GZ: %s -> %s", src, dest))
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dest, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), os.ModePerm); err != nil {
				return err
			}
			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return flattenSingleFolder(dest)
}

// Flatten folder if it contains only a single folder
func flattenSingleFolder(path string) error {
	for {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		// Only one entry and it's a folder
		if len(entries) == 1 && entries[0].IsDir() {
			innerPath := filepath.Join(path, entries[0].Name())
			innerEntries, err := os.ReadDir(innerPath)
			if err != nil {
				return err
			}

			// Move all inner files/folders up
			for _, e := range innerEntries {
				oldPath := filepath.Join(innerPath, e.Name())
				newPath := filepath.Join(path, e.Name())
				if err := os.Rename(oldPath, newPath); err != nil {
					return err
				}
			}

			// Remove the now-empty folder
			if err := os.Remove(innerPath); err != nil {
				return err
			}
		} else {
			break
		}
	}
	return nil
}

// Main function to process downloads
func ProcessDownloads(opts *args.Options) error {
	downloadsPath := filepath.Join(opts.Path, "downloads")
	entries, err := os.ReadDir(downloadsPath)
	if err != nil {
		return fmt.Errorf("failed to read downloads folder: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		folderName := entry.Name()
		folderPath := filepath.Join(downloadsPath, folderName)
		printer.PrintSuccess(fmt.Sprintf("Processing folder: %s", folderName))

		platforms, err := os.ReadDir(folderPath)
		if err != nil {
			return fmt.Errorf("failed to read platform folders in %s: %w", folderName, err)
		}

		for _, platform := range platforms {
			if !platform.IsDir() {
				continue
			}
			platformName := platform.Name()
			platformPath := filepath.Join(folderPath, platformName)
			printer.PrintSuccess(fmt.Sprintf("Processing platform: %s", platformName))

			files, err := os.ReadDir(platformPath)
			if err != nil {
				return fmt.Errorf("failed to read files in %s/%s: %w", folderName, platformName, err)
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}

				srcFile := filepath.Join(platformPath, file.Name())
				destDir := filepath.Join(opts.Path, "bin", folderName, platformName)

				if strings.HasSuffix(file.Name(), ".zip") {
					if err := extractZip(srcFile, destDir); err != nil {
						printer.PrintError(fmt.Sprintf("Failed to extract ZIP: %s, error: %v", srcFile, err))
						return err
					}
				} else if strings.HasSuffix(file.Name(), ".tar.gz") || strings.HasSuffix(file.Name(), ".tgz") {
					if err := extractTarGz(srcFile, destDir); err != nil {
						printer.PrintError(fmt.Sprintf("Failed to extract TAR.GZ: %s, error: %v", srcFile, err))
						return err
					}
				} else {
					printer.PrintSuccess(fmt.Sprintf("Skipping unknown file type: %s", srcFile))
				}
			}
		}
	}

	return nil
}
