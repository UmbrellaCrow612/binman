package extractor

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// CopyToBin copies the content of the first folder with files from downloads into the bin directory
func CopyToBin(binary *shared.Binary, options *args.Options) error {
	baseDownloadDir := filepath.Join(options.Path, "downloads", binary.NAME)
	baseBinDir := filepath.Join(options.Path, "bin", binary.NAME)

	for platform, archMap := range binary.URLS {
		for arch := range archMap {
			downloadArchDir := filepath.Join(baseDownloadDir, platform, arch)

			sourceContentDir, err := findContentDir(downloadArchDir)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return err
			}

			binArchDir := filepath.Join(baseBinDir, platform, arch)
			err = os.MkdirAll(binArchDir, os.ModePerm)
			if err != nil {
				return err
			}

			if err := copyDirContents(sourceContentDir, binArchDir); err != nil {
				return err
			}
		}
	}

	return nil
}

// findContentDir recursively searches a directory to find the deepest subdirectory that
// contains actual files, skipping intermediate, unnecessary folders created by extraction.
func findContentDir(root string) (string, error) {
	currentDir := root

	for {
		// Read directory contents
		entries, err := os.ReadDir(currentDir)
		if err != nil {
			return "", err
		}

		var dirs []fs.DirEntry
		fileCount := 0

		for _, entry := range entries {
			// Skip archives and hidden directories
			if strings.HasPrefix(entry.Name(), ".") || strings.Contains(entry.Name(), ".tar.gz") {
				continue
			}

			if entry.IsDir() {
				dirs = append(dirs, entry)
			} else {
				fileCount++
			}
		}

		// Case 1: The current directory has files. This is the content directory we want.
		if fileCount > 0 {
			return currentDir, nil
		}

		// Case 2: The current directory only contains ONE non-hidden subdirectory and no files.
		// This means we need to descend further.
		if len(dirs) == 1 {
			currentDir = filepath.Join(currentDir, dirs[0].Name())
			continue
		}

		// Case 3: The directory is empty or has multiple subdirectories and no files.
		// If multiple subdirectories exist, we assume this is the top level of the contents.
		return currentDir, nil
	}
}

// copyDirContents copies the contents (files and subdirectories) from src to dst.
func copyDirContents(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectories
			if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
				return err
			}
			if err := copyDirContents(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Copy files with permission preservation
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// copyFile copies a single file from src to dst, preserving the file's permission mode.
func copyFile(src, dst string) error {
	// 1. Get source file info to preserve permissions later
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	// 2. Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 3. Create the destination file with the same permissions
	// os.Create uses 0666, so we'll use os.OpenFile to set initial permissions based on source
	// Note: We use 0755 as a default for new executables if we can't determine the source mode easily.
	// For production, using the source mode is best practice.
	destinationFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode().Perm())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 4. Copy the content efficiently using io.Copy
	if _, err = io.Copy(destinationFile, sourceFile); err != nil {
		return err
	}

	// 5. Ensure the executable permissions are correctly applied to the destination file
	return os.Chmod(dst, info.Mode().Perm())
}
