package extractor

import (
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// CopyToBin copies the first folder with content from downloads into the bin directory
func CopyToBin(binary *shared.Binary, options *args.Options) error {
	baseDownloadDir := filepath.Join(options.Path, "downloads", binary.NAME)
	baseBinDir := filepath.Join(options.Path, "bin", binary.NAME)

	for platform, archMap := range binary.URLS {
		for arch := range archMap {
			downloadArchDir := filepath.Join(baseDownloadDir, platform, arch)
			_, err := os.Stat(downloadArchDir)
			if err != nil {
				return err
			}

			binArchDir := filepath.Join(baseBinDir, platform, arch)
			err = os.MkdirAll(binArchDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
