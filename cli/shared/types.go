package shared

import (
	"errors"
	"fmt"
	"regexp"
)

type Config struct {
	Binaries []Binary `yaml:"binaries"`
}

type Binary struct {
	// Name
	Name string `yaml:"name"`

	// Key is platform, value is URL
	URL map[string]string `yaml:"url"`

	// Key value pair of platform then sha key
	SHA256 map[string]string `yaml:"sha256"`

	// Optional list of keys for the platform downloaded like `windows` then a regular expression pattern to match only specificed files to keep
	Pattern map[string]string `yaml:"pattern"`

	// Another other Key value pair
	Extra map[string]any `yaml:",inline"`
}

// Validate checks if required fields are present for a single Binary
func (b *Binary) Validate() error {
	if b.Name == "" {
		return errors.New("name is required")
	}
	if len(b.URL) == 0 {
		return errors.New("url is required")
	}
	if len(b.SHA256) == 0 {
		return errors.New("sha256 is required")
	}
	return nil
}

// Validate checks the entire Config
func (c *Config) Validate() error {
	if len(c.Binaries) == 0 {
		return errors.New("binaries list is required and cannot be empty")
	}

	for i, b := range c.Binaries {
		if err := b.Validate(); err != nil {
			return fmt.Errorf("binary at index %d: %w", i, err)
		}
	}

	return nil
}

// CompiledPatterns stores compiled regexes per binary and platform binary name → platform → compiled regex
type CompiledPatterns map[string]map[string]*regexp.Regexp

// CompilePatterns validates and compiles all regex patterns for all binaries
func (c *Config) CompilePatterns() (CompiledPatterns, error) {
	result := make(CompiledPatterns)

	for _, b := range c.Binaries {
		if len(b.Pattern) == 0 {
			continue
		}

		result[b.Name] = make(map[string]*regexp.Regexp)
		for platform, pattern := range b.Pattern {
			if pattern == "" {
				continue
			}
			regex, err := regexp.Compile(pattern)
			if err != nil {
				return nil, fmt.Errorf("binary %s, platform %s: invalid pattern %q: %w", b.Name, platform, pattern, err)
			}
			result[b.Name][platform] = regex
		}
	}

	return result, nil
}
