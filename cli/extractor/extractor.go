package extractor

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/console"
	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Extract a specific downloaded asset from ./downloads/packname/platform/arch/example.zip
// To ./extracted/packname/platform/arch/....
// It will only extract folder that have been downloaded from the previous step
// as the previous step contains the filerting logic
func Extract(p *t.Package, o *t.ArgOptions) error {
	extractedDir, err := filepath.Abs(filepath.Join(o.BasePath, "extracted"))
	if err != nil {
		return err
	}

	packnameDir := filepath.Join(extractedDir, p.Name)

	for platform, archMap := range p.Platforms {
		platformDir := filepath.Join(packnameDir, platform)

		for arch, asset := range archMap {
			// Points to where we will extract the content to ./extracted/packname/platform/arch/...
			archDir := filepath.Join(platformDir, arch)

			fileName := strings.TrimSpace(path.Base(asset.URL))
			if fileName == "" || fileName == "." {
				return errors.New("URL does not point to a file: " + asset.URL)
			}

			// points to where it should be downloaded ./downloads/packname/platform/arch/example.zip from URL
			downloadedDir, err := filepath.Abs(filepath.Join(o.BasePath, "downloads", p.Name, platform, arch, fileName))
			if err != nil {
				return err
			}

			if !pathExists(downloadedDir) {
				console.LogInfo("Skipping extract of " + downloadedDir)
				continue
			}

			extract(downloadedDir, archDir)
		}
	}

	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Interal impl that takes a from path like ./downloads/packname/platform/arch/example.zip
// To ./extracted/packname/platform/arch/...
func extract(fromPath string, toPath string) error {
	ext := path.Ext(fromPath)
	if ext == "" {
		return errors.New("Downloaded path does not have a extension " + fromPath)
	}

	switch ext {
	case "":
	default:
		return errors.New("File extraction failed as extension not supported " + ext)

	}
	return nil
}
