package extractor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

func CopyToBin(bin *shared.Binary, options *args.Options) error {
	baseDownloadDir := filepath.Join(options.Path, "downloads", bin.NAME)
	binDir := filepath.Join(options.Path, "bin", bin.NAME)

	if _, err := os.Stat(baseDownloadDir); os.IsNotExist(err) {
		return fmt.Errorf("download folder %s not found", baseDownloadDir)
	}

	for platform, archAndUrlMap := range bin.URLS {
		for arch := range archAndUrlMap {
			finalDownloadDir := filepath.Join(baseDownloadDir, platform, arch)
			_, err := os.Stat(finalDownloadDir)
			if err != nil {
				return fmt.Errorf("download folder %s not found", finalDownloadDir)
			}

			if err := removeArchives(finalDownloadDir); err != nil {
				return err
			}
			if err := flattenFolder(finalDownloadDir); err != nil {
				return err
			}
			if err := removeEmptyDirs(finalDownloadDir); err != nil {
				return err
			}

			finalBinDir := filepath.Join(binDir, platform, arch)
			if err := os.MkdirAll(finalBinDir, 0755); err != nil {
				return fmt.Errorf("failed to create bin directory %s: %w", finalBinDir, err)
			}

			files, err := os.ReadDir(finalDownloadDir)
			if err != nil {
				return fmt.Errorf("failed to read files in %s: %w", finalDownloadDir, err)
			}

			for _, file := range files {
				srcPath := filepath.Join(finalDownloadDir, file.Name())
				dstPath := filepath.Join(finalBinDir, file.Name())

				counter := 1
				for {
					if _, err := os.Stat(dstPath); os.IsNotExist(err) {
						break
					}
					ext := filepath.Ext(file.Name())
					name := file.Name()[0 : len(file.Name())-len(ext)]
					dstPath = filepath.Join(finalBinDir, fmt.Sprintf("%s_%d%s", name, counter, ext))
					counter++
				}

				if err := os.Rename(srcPath, dstPath); err != nil {
					return fmt.Errorf("failed to move file %s to %s: %w", srcPath, dstPath, err)
				}
			}
		}
	}

	return nil
}

// removeArchives deletes all archive files in the directory and subdirectories
func removeArchives(root string) error {
	archiveExts := []string{".zip", ".tar", ".tar.gz", ".tgz", ".tar.bz2", ".tar.xz"}
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		lowerName := strings.ToLower(info.Name())
		for _, ext := range archiveExts {
			if strings.HasSuffix(lowerName, ext) {
				if err := os.Remove(path); err != nil {
					return fmt.Errorf("failed to remove archive %s: %w", path, err)
				}
				break
			}
		}

		return nil
	})
}

// flattenFolder moves all files from subdirectories to the root
func flattenFolder(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root || info.IsDir() {
			return nil
		}

		destPath := filepath.Join(root, info.Name())

		// Handle duplicates
		counter := 1
		for {
			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				break
			}
			ext := filepath.Ext(info.Name())
			name := info.Name()[0 : len(info.Name())-len(ext)]
			destPath = filepath.Join(root, fmt.Sprintf("%s_%d%s", name, counter, ext))
			counter++
		}

		return os.Rename(path, destPath)
	})
}

// removeEmptyDirs deletes empty directories recursively
func removeEmptyDirs(root string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(root, entry.Name())
			if err := removeEmptyDirs(subDir); err != nil {
				return err
			}
		}
	}

	entries, err = os.ReadDir(root)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return os.Remove(root)
	}

	return nil
}
