package transfer

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/UmbrellaCrow612/binman/console"
	"github.com/UmbrellaCrow612/binman/global"
	"github.com/UmbrellaCrow612/binman/t"
)

// Transfer the extracted package content to the final destination bin folder using the pattern to only transfer
// the required files needed
// it will look at BinManExtractedPath/packname/platform/arch/... if it exists it will then move it to
// BinManBinPath/packname/platform/arch/....
func ToBin(p *t.Package, o *t.ArgOptions) error {
	binDir := global.BinmanBinPath

	packDir := filepath.Join(binDir, p.Name)

	for platform, archMap := range p.Platforms {
		platformDir := filepath.Join(packDir, platform)
		for arch, asset := range archMap {
			archDir := filepath.Join(platformDir, arch)

			// points to the extracted folder it should have bee extracted to ./extracted/packname/platform/arch/...
			// if it does not exist skip transfer step as it was filtered out
			extractedDir := global.BinmanExtractedPath
			if !pathExists(extractedDir) {
				console.LogInfo("Skipping transerfing of " + extractedDir)
				continue
			}

			reg, err := regexp.Compile(asset.Pattern)
			if err != nil {
				return err
			}

			console.LogInfo("Transfering extracted content to " + archDir)
			err = toBin(extractedDir, archDir, reg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Transfer the files from the from path to the folder of the to path only transfering files
// that match the regular expression
func toBin(fromPath string, toPath string, reg *regexp.Regexp) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("from path is not a directory: " + fromPath)
	}

	if err := os.MkdirAll(toPath, 0755); err != nil {
		return err
	}

	return filepath.WalkDir(fromPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !reg.MatchString(d.Name()) {
			return nil
		}

		destPath := filepath.Join(toPath, d.Name())

		if err := os.Rename(path, destPath); err != nil {
			return err
		}

		return nil
	})
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
