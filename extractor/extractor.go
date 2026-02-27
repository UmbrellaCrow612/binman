package extractor

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/console"
	"github.com/UmbrellaCrow612/binman/global"
	"github.com/UmbrellaCrow612/binman/t"
)

// Extract a specific downloaded asset from BinManDownloadPath/example.zip
// To BinManExtractedPath/extracted/packname/platform/arch/....
// It will only extract folder that have been downloaded from the previous step
// as the previous step contains the filerting logic
func Extract(p *t.Package, o *t.ArgOptions) error {
	extractedDir := global.BinmanExtractedPath

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

			// points to where it should be downloaded BinmanDownloadPath/packname/platform/arch/example.zip from URL
			downloadedDir := filepath.Join(global.BinmanDownloadPath, p.Name, platform, arch, fileName)

			if !pathExists(downloadedDir) {
				console.LogInfo("Skipping extract of " + downloadedDir)
				continue
			}

			console.LogInfo("Extracting content to " + archDir)
			if err := extract(downloadedDir, archDir); err != nil {
				return err
			}
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
	case ".zip":
		err := extractZip(fromPath, toPath)
		return err
	case ".tar":
		err := ExtractTar(fromPath, toPath)
		return err
	case ".gz":
		err := ExtractTarGz(fromPath, toPath)
		return err
	default:
		return errors.New("File extraction failed as extension not supported " + ext)

	}
}
