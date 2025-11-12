package fetch

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"net/http"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/yml"
)

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
			printer.PrintError(fmt.Sprintf("Failed to fetch %s URL for %s: %v", osKey, bin.Name, err))
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			printer.PrintError(fmt.Sprintf("Error fetching %s URL for %s (status: %d)", osKey, bin.Name, resp.StatusCode))
			resp.Body.Close()
			continue
		}

		// Determine file extension from URL
		filename := filepath.Base(url)
		filePath := filepath.Join(baseDir, filename)

		out, err := os.Create(filePath)
		if err != nil {
			printer.PrintError(fmt.Sprintf("Failed to create file %s: %v", filePath, err))
			resp.Body.Close()
			continue
		}

		_, err = io.Copy(out, resp.Body)
		resp.Body.Close()
		out.Close()
		if err != nil {
			printer.PrintError(fmt.Sprintf("Failed to write file %s: %v", filePath, err))
			continue
		}

		printer.PrintSuccess(fmt.Sprintf("Downloaded %s -> %s", osKey, filePath))
	}
}