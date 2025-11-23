package fetch

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

// Fetches the binary urls into path/downloads
// downloads all of them into the convention
// opts.PATH/downloads/ripgrep/linux/x86_64/ripgrep.zip
func FetchAndStoreBinary(bin *shared.Binary, opts *args.Options) error {

	// Base dir becomes example downloads/ripgrep
	baseDir := filepath.Join(opts.Path, "downloads", bin.NAME)
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create download directory: %w", err)
	}

	for platform, architectureAndUrl := range bin.URLS {
		if len(opts.SpecificPlatformBuilds) > 0 &&
			!slices.Contains(opts.SpecificPlatformBuilds, platform) {
			printer.PrintSuccess("Skipping fetch " + platform)
			continue
		}

		for architecture, url := range architectureAndUrl {
			if len(opts.SpecificArchBuilds) > 0 && !slices.Contains(opts.SpecificArchBuilds, architecture) {
				printer.PrintSuccess("Skipping fetch " + architecture)
				continue
			}

			printer.PrintSuccess("Fetching " + url)

			// Example: downloads/ripgrep/linux/x86_64
			finalDir := filepath.Join(baseDir, platform, architecture)
			if err := os.MkdirAll(finalDir, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create download directory: %w", err)
			}

			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("failed to fetch %s: %w", url, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to fetch %s: status %s", url, resp.Status)
			}

			// Determine filename from the URL
			parts := strings.Split(url, "/")
			fileName := parts[len(parts)-1]
			filePath := filepath.Join(finalDir, fileName)

			out, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", filePath, err)
			}
			defer out.Close()

			if _, err := io.Copy(out, resp.Body); err != nil {
				return fmt.Errorf("failed to write file %s: %w", filePath, err)
			}

			expectedSHA, ok := bin.SHA256[platform][architecture]
			if !ok {
				return fmt.Errorf("no SHA256 provided for %s/%s", platform, architecture)
			}

			if err := VerifySHA256(filePath, expectedSHA); err != nil {
				return err
			}
		}
	}

	return nil
}

// Helper function to check SHA256 of a file
func VerifySHA256(filePath, expectedSHA string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return fmt.Errorf("failed to hash file %s: %w", filePath, err)
	}

	actualSHA := hex.EncodeToString(hasher.Sum(nil))
	if actualSHA != expectedSHA {
		return fmt.Errorf(
			"SHA256 mismatch for %s: expected %s, got %s",
			filePath, expectedSHA, actualSHA,
		)
	}

	printer.PrintSuccess("SHA256 verified for " + filePath)
	return nil
}
