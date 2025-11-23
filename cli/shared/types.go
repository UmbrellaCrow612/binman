package shared

import (
	"errors"
	"fmt"
	"regexp"
)

// Represents the binman.yml
type Config struct {
	Binaries []Binary `yaml:"binaries"`
}

// Validate checks that the config has at least one binary and each binary is valid
func (c *Config) Validate() error {
	if len(c.Binaries) == 0 {
		return errors.New("config must contain at least one binary")
	}

	for i, bin := range c.Binaries {
		if err := bin.Validate(); err != nil {
			return fmt.Errorf("binary at index %d validation failed: %w", i, err)
		}
	}

	return nil
}

// Represents a specific binary defined to be downloaded
type Binary struct {
	// Represents the name field of a binary yml
	NAME string `yaml:"name"`

	// URLS represents a mapping of platform -> architecture -> download URL.
	//
	// Example structure (YAML):
	//
	// urls:
	//   linux:
	//     x86_64: https://example.com/rg-x86_64.tar.gz
	//     arm64:  https://example.com/rg-arm64.tar.gz
	//
	// Meaning:
	//   platform → architectures → URL for each architecture
	//
	URLS map[string]map[string]string `yaml:"urls"`

	// Represents a mapping of platform -> architecture -> SHA256 checksum.
	//
	// Example structure (YAML):
	//
	// sha256:
	//   linux:
	//     x86_64: 1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599
	//     arm64: <real checksum here>
	//   darwin:
	//     x86_64: 64811cb24e77cac3057d6c40b63ac9becf9082eedd54ca411b475b755d334882
	//     arm64: <real checksum here>
	//   windows:
	//     x86_64: 124510b94b6baa3380d051fdf4650eaa80a302c876d611e9dba0b2e18d87493a
	//     arm64: <real checksum here>
	//
	SHA256 map[string]map[string]string `yaml:"sha256"`

	// PATTERNS represents the executable name pattern (regex) for each platform -> architecture.
	//
	// Example structure (YAML):
	//
	// patterns:
	//   linux:
	//     x86_64: "^rg$"
	//     arm64: "^rg$"
	//   darwin:
	//     x86_64: "^rg$"
	//     arm64: "^rg$"
	//   windows:
	//     x86_64: "^rg\\.exe$"
	//     arm64: "^rg\\.exe$"
	//
	// Meaning:
	//   platform → architectures → executable name pattern (regex)
	//
	PATTERNS map[string]map[string]string `yaml:"patterns"`

	// EXTRA captures any other key-value pairs in the YAML that are not explicitly mapped.
	// For example, version: "1.2.3" would go here.
	EXTRA map[string]interface{} `yaml:",inline"`
}

// Validate checks that the binary has required fields
func (b *Binary) Validate() error {
	if b.NAME == "" {
		return errors.New("binary name cannot be empty")
	}

	if len(b.URLS) == 0 {
		return fmt.Errorf("binary '%s' must have at least one URL defined", b.NAME)
	}

	if len(b.SHA256) == 0 {
		return fmt.Errorf("binary '%s' must have at least one SHA256 checksum defined", b.NAME)
	}

	for platform, arches := range b.PATTERNS {
		for arch, pattern := range arches {
			if _, err := regexp.Compile(pattern); err != nil {
				return fmt.Errorf("invalid pattern for binary '%s', platform '%s', architecture '%s': '%s' -> %w", b.NAME, platform, arch, pattern, err)
			}
		}
	}

	return nil
}

// CompilePatternsMap compiles all patterns and returns a nested map of platform -> architecture -> *regexp.Regexp.
// If any pattern fails, it returns an error indicating which platform/architecture failed.
func (b *Binary) CompilePatternsMap() (map[string]map[string]*regexp.Regexp, error) {
	compiled := make(map[string]map[string]*regexp.Regexp)

	for platform, arches := range b.PATTERNS {
		if compiled[platform] == nil {
			compiled[platform] = make(map[string]*regexp.Regexp)
		}

		for arch, pattern := range arches {
			re, err := regexp.Compile(pattern)
			if err != nil {
				return nil, fmt.Errorf("invalid pattern for platform '%s', architecture '%s': '%s' -> %w", platform, arch, pattern, err)
			}
			compiled[platform][arch] = re
		}
	}

	return compiled, nil
}
