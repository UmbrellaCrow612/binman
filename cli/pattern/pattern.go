package pattern

import (
	"os"
	"path/filepath"
	"slices"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// Cleans a specific binary's bin with patterns defined for it
func CleanWithPattern(bin *shared.Binary, options *args.Options) error {
	baseBinDir := filepath.Join(options.Path, "bin", bin.NAME)

	compliedRegexMap, err := bin.CompilePatternsMap()
	if err != nil {
		return err
	}

	for platform, archAndUrl := range bin.URLS {
		if len(options.SpecificPlatformBuilds) > 0 &&
			!slices.Contains(options.SpecificPlatformBuilds, platform) {
			continue
		}

		for arch := range archAndUrl {
			if len(options.SpecificArchBuilds) > 0 && !slices.Contains(options.SpecificArchBuilds, arch) {
				continue
			}

			regex, ok := compliedRegexMap[platform][arch]
			if !ok {
				printer.PrintWarning("Platform " + platform + " " + arch + " has no defined pattern, skipping")
				continue
			}

			finalBinDir := filepath.Join(baseBinDir, platform, arch)
			info, err := os.Stat(finalBinDir)
			if os.IsNotExist(err) || !info.IsDir() {
				continue
			} else if err != nil {
				return err
			}

			// Remove files not matching regex
			err = filepath.WalkDir(finalBinDir, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() && !regex.MatchString(d.Name()) {
					return os.Remove(path)
				}
				return nil
			})
			if err != nil {
				return err
			}

			// Remove empty directories
			err = filepath.Walk(finalBinDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					entries, readErr := os.ReadDir(path)
					if readErr != nil {
						return readErr
					}
					if len(entries) == 0 {
						return os.Remove(path)
					}
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
