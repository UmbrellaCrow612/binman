package fetch

import (
	"path/filepath"
	"slices"

	"github.com/UmbrellaCrow612/binman/cli/console"
	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Fetch a package into ./downloads/packname/platform/arch/...
// Using the options defined
func Get(p *t.Package, o *t.ArgOptions) error {
	downloadDir, err := filepath.Abs(filepath.Join(o.BasePath, "downloads"))
	if err != nil {
		return err
	}

	console.LogInfo("Download package " + p.Name)

	packageDir := filepath.Join(downloadDir, p.Name)

	for p, archMap := range p.Platforms {
		if len(o.Platforms) > 0 && !slices.Contains(o.Platforms, p) {
			console.LogInfo("Skipping platform " + p)
			continue
		}
		platformDir := filepath.Join(packageDir, p)

		for arch := range archMap {
			if len(o.Architectures) > 0 && !slices.Contains(o.Architectures, arch) {
				console.LogInfo("Skipping architectures " + arch)
				continue
			}
			finalDownloadDir := filepath.Join(platformDir, arch)

			console.LogInfo("Downloading content to " + finalDownloadDir)
		}
	}

	return nil
}
