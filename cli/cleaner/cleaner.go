package cleaner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// clean deletes the bin and downloads folders inside the provided path
func Clean(opts *args.Options) {
	folders := []string{
		filepath.Join(opts.Path, "bin"),
		filepath.Join(opts.Path, "downloads"),
	}

	for _, folder := range folders {
		err := os.RemoveAll(folder)
		if err != nil {
			printer.PrintError("Failed to remove folder: " + folder + " - " + err.Error())
			continue
		}
		printer.PrintSuccess("Successfully removed folder: " + folder)
	}
}

// CleanDownloads deletes the downloads folder inside the provided path
func CleanDownloads(opts *args.Options) {
	downloadsPath := filepath.Join(opts.Path, "downloads")

	err := os.RemoveAll(downloadsPath)
	if err != nil {
		printer.PrintError("Failed to remove downloads folder: " + downloadsPath + " - " + err.Error())
		return
	}

	printer.PrintSuccess("Successfully removed downloads folder: " + downloadsPath)
}

// Cleans downoaded bin folder based on pattern defined
func CleanBin(options *args.Options, compiledPatterns shared.CompiledPatterns) error {
	binPath := filepath.Join(options.Path, "bin")
	printer.PrintSuccess(fmt.Sprintf("Cleaning bin folder: %s", binPath))

	// Iterate over all binaries in compiledPatterns
	for binaryName, platformPatterns := range compiledPatterns {
		binaryFolder := filepath.Join(binPath, binaryName)

		info, err := os.Stat(binaryFolder)
		if err != nil {
			if os.IsNotExist(err) {
				printer.PrintError(fmt.Sprintf("Binary folder does not exist, skipping: %s", binaryFolder))
				continue
			}
			return fmt.Errorf("failed to access binary folder %s: %w", binaryFolder, err)
		}
		if !info.IsDir() {
			printer.PrintError(fmt.Sprintf("Expected binary folder is not a directory, skipping: %s", binaryFolder))
			continue
		}

		printer.PrintSuccess(fmt.Sprintf("Processing binary: %s", binaryName))

		// Iterate over each platform folder
		for platform, regex := range platformPatterns {
			if options.SpecificPlatformBuild != "" && options.SpecificPlatformBuild != platform {
				printer.PrintSuccess(fmt.Sprintf("Skipping platform %s (SpecificPlatformBuild=%s)", platform, options.SpecificPlatformBuild))
				continue
			}

			platformFolder := filepath.Join(binaryFolder, platform)
			info, err := os.Stat(platformFolder)
			if err != nil {
				if os.IsNotExist(err) {
					printer.PrintError(fmt.Sprintf("Platform folder does not exist, skipping: %s", platformFolder))
					continue
				}
				return fmt.Errorf("failed to access platform folder %s: %w", platformFolder, err)
			}
			if !info.IsDir() {
				printer.PrintError(fmt.Sprintf("Expected platform folder is not a directory, skipping: %s", platformFolder))
				continue
			}

			printer.PrintSuccess(fmt.Sprintf("Cleaning platform folder: %s", platformFolder))

			// Walk through files in the platform folder
			err = filepath.Walk(platformFolder, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				// Remove files that do not match the regex
				if !regex.MatchString(info.Name()) {
					printer.PrintSuccess(fmt.Sprintf("Removing file: %s", path))
					if err := os.Remove(path); err != nil {
						return fmt.Errorf("failed to remove file %s: %w", path, err)
					}
				} else {
					printer.PrintSuccess(fmt.Sprintf("Keeping file: %s", path))
				}
				return nil
			})
			if err != nil {
				return fmt.Errorf("error cleaning platform folder %s: %w", platformFolder, err)
			}
		}
	}

	printer.PrintSuccess("Bin folder cleaned successfully.")
	return nil
}
