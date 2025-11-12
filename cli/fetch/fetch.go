package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/yml"
)

// FetchAndStoreBinary downloads all URLs of a binary and stores them under opts.Path/downloads
// Fail-fast: exits on first error
func FetchAndStoreBinary(bin *yml.Binary, opts *args.Options) {
	baseDir := filepath.Join(opts.Path, "downloads", bin.Name)
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		printer.ExitError(fmt.Sprintf("Failed to create download directory: %v", err))
	}

	client := http.Client{Timeout: 20 * time.Second}

	for osKey, url := range bin.URL {
		printer.PrintSuccess(fmt.Sprintf("Fetching %s -> %s", osKey, url))

		resp, err := client.Get(url)
		if err != nil {
			printer.ExitError(fmt.Sprintf("Failed to fetch %s URL for %s: %v", osKey, bin.Name, err))
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			printer.ExitError(fmt.Sprintf("Error fetching %s URL for %s (status: %d)", osKey, bin.Name, resp.StatusCode))
		}

		// Determine file path
		filename := filepath.Base(url)
		filePath := filepath.Join(baseDir, filename)

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

		printer.PrintSuccess(fmt.Sprintf("Downloaded %s -> %s", osKey, filePath))
	}
}