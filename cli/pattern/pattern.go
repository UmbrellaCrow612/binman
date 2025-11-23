package pattern

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// Cleans a specific binarys bin with patterns defined for it
func CleanWithPattern(bin *shared.Binary, options *args.Options) error {
	baseBinDir := filepath.Join(options.Path, "bin", bin.NAME)
	_, err := os.Stat(baseBinDir)
	if err != nil {
		return err
	}

	compiledRegex, err := bin.CompilePatternsMap()
	if err != nil {
		return err
	}

	for platform, archAndRegex := range bin.PATTERNS {
		for arch := range archAndRegex {
			regex, ok := compiledRegex[platform][arch]
			if !ok {
				return fmt.Errorf("no pattern found for platform '%s' and architecture '%s'", platform, arch)
			}

			finalBinDir := filepath.Join(baseBinDir, platform, arch)
			_, err := os.Stat(finalBinDir)
			if err != nil {
				return err
			}

			printer.PrintSuccess("Cleaning " + finalBinDir)

			// Walk recursively
			err = filepath.WalkDir(finalBinDir, func(path string, d os.DirEntry, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}

				// Skip directories
				if d.IsDir() {
					return nil
				}

				// If file does not match pattern, remove it
				if !regex.MatchString(d.Name()) {
					if err := os.Remove(path); err != nil {
						return fmt.Errorf("failed to remove file '%s': %w", path, err)
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
