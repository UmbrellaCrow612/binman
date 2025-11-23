package pattern

import (
	"os"
	"path/filepath"
	"slices"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// Cleans a specific binarys bin with patterns defined for it
func CleanWithPattern(bin *shared.Binary, options *args.Options) error {
	baseBinDir := filepath.Join(options.Path, "bin", bin.NAME)
	if _, err := os.Stat(baseBinDir); err != nil {
		return err
	}

	compliedRegexMap, err := bin.CompilePatternsMap()
	if err != nil {
		return err
	}

	for platform, archAndUrl := range bin.URLS {
		if len(options.SpecificPlatformBuilds) > 0 &&
			!slices.Contains(options.SpecificPlatformBuilds, platform) {
			printer.PrintSuccess("Skipping pattern cleaning" + platform)
			continue
		}

		for arch := range archAndUrl {
			if len(options.SpecificArchBuilds) > 0 && !slices.Contains(options.SpecificArchBuilds, architecture) {
				printer.PrintSuccess("Skipping pattern cleaning" + arch)
				continue
			}

			regex, ok := compliedRegexMap[platform][arch]
			if !ok {
				printer.PrintWarning("Platform " + platform + " " + arch + "has not defined a pattern so skipping")
				continue
			}

			finalBinDir := filepath.Join(baseBinDir, platform, arch)
			_, err := os.Stat(finalBinDir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
