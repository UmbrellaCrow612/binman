package fetch

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

func FetchAndStoreBinary(bin *shared.Binary, opts *args.Options) {
	baseDir := filepath.Join(opts.Path, "downloads", bin.Name)
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		printer.ExitError(fmt.Sprintf("Failed to create download directory: %v", err))
	}

	client := http.Client{Timeout: 20 * time.Second}

	if opts.SpecificPlatformBuild != "" {
		if _, ok := bin.URL[opts.SpecificPlatformBuild]; !ok {
			printer.ExitError(fmt.Sprintf(
				"Binary '%s' does not define a URL for platform '%s'",
				bin.Name,
				opts.SpecificPlatformBuild,
			))
		}
	}

	for platform, url := range bin.URL {
		if opts.SpecificPlatformBuild != "" && platform != opts.SpecificPlatformBuild {
			continue
		}

		printer.PrintSuccess(fmt.Sprintf("Fetching %s -> %s", platform, url))

		resp, err := client.Get(url)
		if err != nil {
			printer.ExitError(fmt.Sprintf("Failed to fetch %s URL for %s: %v", platform, bin.Name, err))
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			printer.ExitError(fmt.Sprintf("Error fetching %s URL for %s (status: %d)", platform, bin.Name, resp.StatusCode))
		}

		platformDir := filepath.Join(baseDir, platform)
		if err := os.MkdirAll(platformDir, os.ModePerm); err != nil {
			printer.ExitError(fmt.Sprintf("Failed to create platform directory %s: %v", platformDir, err))
		}

		filename := filepath.Base(url)
		filePath := filepath.Join(platformDir, filename)

		out, err := os.Create(filePath)
		if err != nil {
			printer.ExitError(fmt.Sprintf("Failed to create file %s: %v", filePath, err))
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			out.Close()
			printer.ExitError(fmt.Sprintf("Failed to write file %s: %v", filePath, err))
		}

		if err := out.Close(); err != nil {
			printer.ExitError(fmt.Sprintf("Failed to close file %s: %v", filePath, err))
		}

		printer.PrintSuccess(fmt.Sprintf("Downloaded %s -> %s", platform, filePath))

		expectedSHA, ok := bin.SHA256[platform]
		if !ok || expectedSHA == "" {
			printer.ExitError(fmt.Sprintf("No SHA256 provided for %s on platform %s", bin.Name, platform))
		}

		if err := verifySHA256(filePath, expectedSHA); err != nil {
			printer.ExitError(fmt.Sprintf("SHA256 verification failed for %s: %v", filePath, err))
		}

		printer.PrintSuccess(fmt.Sprintf("SHA256 verified for %s", filePath))
	}
}

func verifySHA256(filePath, expected string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return fmt.Errorf("cannot compute hash: %v", err)
	}

	actual := hex.EncodeToString(hash.Sum(nil))
	if actual != expected {
		return fmt.Errorf("hash mismatch: expected %s, got %s", expected, actual)
	}
	return nil
}
