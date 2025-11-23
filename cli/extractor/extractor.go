package extractor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/args"
)

// getArchiveFolders walks through options.PATH/downloads and collects all archive files
// Returns a slice of archive paths or an error if something goes wrong
func getArchiveFolders(options *args.Options) ([]string, error) {
	if options == nil || options.Path == "" {
		return nil, fmt.Errorf("invalid options or PATH")
	}

	downloadsPath := filepath.Join(options.Path, "downloads")
	var archives []string

	err := filepath.WalkDir(downloadsPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %w", path, err)
		}

		if d.IsDir() {
			return nil
		}

		lower := strings.ToLower(d.Name())
		if strings.HasSuffix(lower, ".zip") ||
			strings.HasSuffix(lower, ".tar") ||
			strings.HasSuffix(lower, ".tar.gz") ||
			strings.HasSuffix(lower, ".tgz") {
			archives = append(archives, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return archives, nil
}

func Extract(options *args.Options) error {
	matches, err := getArchiveFolders(options)
	if err != nil {
		return err
	}

	for _, m := range matches {
		lower := strings.ToLower(m)
		ext := filepath.Ext(lower)

		switch ext {
		case ".zip":
			destDir := strings.TrimSuffix(m, ext)
			if err := ExtractZip(m, destDir); err != nil {
				return fmt.Errorf("failed to extract zip %s: %w", m, err)
			}

		case ".tar":
			destDir := strings.TrimSuffix(m, ext)
			if err := ExtractTar(m, destDir); err != nil {
				return fmt.Errorf("failed to extract tar %s: %w", m, err)
			}

		case ".gz":
			if strings.HasSuffix(lower, ".tar.gz") || strings.HasSuffix(lower, ".tgz") {
				destDir := strings.TrimSuffix(strings.TrimSuffix(m, ".tar.gz"), ".tgz")
				if err := ExtractTarGz(m, destDir); err != nil {
					return fmt.Errorf("failed to extract tar.gz %s: %w", m, err)
				}
			} else {
				destDir := strings.TrimSuffix(m, ext)
				if err := ExtractGz(m, destDir); err != nil {
					return fmt.Errorf("failed to extract gz %s: %w", m, err)
				}
			}

		default:
			return fmt.Errorf("unsupported archive type: %s\n", m)
		}
	}

	return nil
}