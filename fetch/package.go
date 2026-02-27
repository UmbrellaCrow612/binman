package fetch

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/console"
	"github.com/UmbrellaCrow612/binman/t"
)

// Fetch a package into ./downloads/packname/platform/arch/...
// Using the options defined
func Get(p *t.Package, o *t.ArgOptions) error {
	if len(o.Packages) > 0 && !slices.Contains(o.Packages, p.Name) {
		console.LogInfo("Skipping downloading package " + p.Name)
		return nil
	}

	downloadDir, err := filepath.Abs(filepath.Join(o.BasePath, "downloads"))
	if err != nil {
		return err
	}

	console.LogInfo("Downloading package: " + p.Name)

	packageDir := filepath.Join(downloadDir, p.Name)

	for p, archMap := range p.Platforms {
		if len(o.Platforms) > 0 && !slices.Contains(o.Platforms, p) {
			console.LogInfo("Skipping platform " + p)
			continue
		}
		platformDir := filepath.Join(packageDir, p)

		console.LogInfo("Downloading platform: " + p)

		for arch, asset := range archMap {
			if len(o.Architectures) > 0 && !slices.Contains(o.Architectures, arch) {
				console.LogInfo("Skipping architectures: " + arch)
				continue
			}
			finalDownloadDir := filepath.Join(platformDir, arch)

			console.LogInfo("Downloading architecture: " + arch)

			err := get(&asset, finalDownloadDir, asset.SHA256)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Download the asset at the given dir
func get(asset *t.Asset, downloadPath string, sha256 string) error {
	fileName := path.Base(strings.Split(asset.URL, "?")[0])
	if strings.TrimSpace(fileName) == "" || fileName == "." {
		return errors.New("URL does not point to a file: " + asset.URL)
	}

	finalDownloadPath := filepath.Join(downloadPath, fileName)

	console.LogInfo("Downloading content to: " + finalDownloadPath)

	resp, err := http.Get(asset.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s, status: %s", asset.URL, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(finalDownloadPath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(finalDownloadPath, data, 0644)
	if err != nil {
		return err
	}

	sha, err := sha256OfFile(finalDownloadPath)
	if err != nil {
		return err
	}

	if sha != sha256 {
		return errors.New("SHA256 " + sha + " does not match " + sha256)
	}

	console.LogInfo("SHA256 verified " + sha256)
	console.LogInfo("Downloaded successfully: " + finalDownloadPath)
	return nil
}

func sha256OfFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
