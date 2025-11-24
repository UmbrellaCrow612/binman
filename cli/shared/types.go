package shared

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/UmbrellaCrow612/binman/cli/args"
)

// Represents the binman.yml
type Config struct {
	Binaries []Binary `yaml:"binaries"`
}

// Validate checks that the config has at least one binary and each binary is valid
func (c *Config) validate() error {
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

	validPlatforms := map[string]bool{
		"linux":   true,
		"windows": true,
		"darwin":  true,
	}

	validArchs := map[string]bool{
		"x86_64":  true,
		"amd64":   true,
		"arm64":   true,
		"aarch64": true,
		"armv7":   true,
		"armv6":   true,
		"i386":    true,
		"386":     true,
	}

	// Convert valid lists to strings for readable error messages
	validPlatformList := make([]string, 0, len(validPlatforms))
	for p := range validPlatforms {
		validPlatformList = append(validPlatformList, p)
	}

	validArchList := make([]string, 0, len(validArchs))
	for a := range validArchs {
		validArchList = append(validArchList, a)
	}

	// Validate all platforms + arches in URLS
	for platform, archMap := range b.URLS {

		if !validPlatforms[platform] {
			return fmt.Errorf(
				"binary '%s' defines invalid platform '%s'. valid platforms: %v",
				b.NAME, platform, validPlatformList,
			)
		}

		if len(archMap) == 0 {
			return fmt.Errorf(
				"binary '%s' platform '%s' must define at least one architecture",
				b.NAME, platform,
			)
		}

		for arch := range archMap {
			if !validArchs[arch] {
				return fmt.Errorf(
					"binary '%s' defines invalid architecture '%s' under platform '%s'. valid architectures: %v",
					b.NAME, arch, platform, validArchList,
				)
			}
		}
	}

	// --- everything below is unchanged ---

	if len(b.URLS) == 0 {
		return fmt.Errorf("binary '%s' must define urls", b.NAME)
	}

	if len(b.SHA256) == 0 {
		return fmt.Errorf("binary '%s' must define sha256", b.NAME)
	}

	for platform, archURLs := range b.URLS {
		shaArchMap, ok := b.SHA256[platform]
		if !ok {
			return fmt.Errorf("binary '%s' missing sha256 definitions for platform '%s'", b.NAME, platform)
		}

		for arch := range archURLs {
			if _, ok := shaArchMap[arch]; !ok {
				return fmt.Errorf(
					"binary '%s' missing sha256 for platform '%s', architecture '%s'",
					b.NAME, platform, arch,
				)
			}
		}
	}

	for platform, shaArchMap := range b.SHA256 {
		if _, ok := b.URLS[platform]; !ok {
			return fmt.Errorf("binary '%s' defines sha256 for platform '%s' but missing in urls", b.NAME, platform)
		}

		for arch := range shaArchMap {
			if _, ok := b.URLS[platform][arch]; !ok {
				return fmt.Errorf(
					"binary '%s' defines sha256 for platform '%s', architecture '%s' but missing in urls",
					b.NAME, platform, arch,
				)
			}
		}
	}

	// Validate patterns (regex)
	for platform, arches := range b.PATTERNS {
		if !validPlatforms[platform] {
			return fmt.Errorf(
				"patterns: invalid platform '%s'. valid platforms: %v",
				platform, validPlatformList,
			)
		}

		for arch, pattern := range arches {
			if !validArchs[arch] {
				return fmt.Errorf(
					"patterns: invalid architecture '%s' for platform '%s'. valid architectures: %v",
					arch, platform, validArchList,
				)
			}

			if _, err := regexp.Compile(pattern); err != nil {
				return fmt.Errorf(
					"invalid pattern for binary '%s', platform '%s', architecture '%s': '%s' -> %w",
					b.NAME, platform, arch, pattern, err,
				)
			}
		}
	}

	return nil
}

func (c *Config) ValidateWithOptions(opts *args.Options) error {
	if err := c.validate(); err != nil {
		return err
	}

	if len(opts.SpecificPlatformBuilds) > 0 {
		for _, bin := range c.Binaries {
			for _, platform := range opts.SpecificPlatformBuilds {

				// Check platform exists in URLs
				if _, ok := bin.URLS[platform]; !ok {
					return fmt.Errorf(
						"binary '%s' does not define platform '%s' in urls",
						bin.NAME, platform,
					)
				}

				// Check platform exists in SHA256
				if _, ok := bin.SHA256[platform]; !ok {
					return fmt.Errorf(
						"binary '%s' does not define platform '%s' in sha256",
						bin.NAME, platform,
					)
				}
			}
		}
	}

	if len(opts.SpecificArchBuilds) > 0 {
		for _, bin := range c.Binaries {

			targetPlatforms := opts.SpecificPlatformBuilds
			if len(targetPlatforms) == 0 {
				for p := range bin.URLS {
					targetPlatforms = append(targetPlatforms, p)
				}
			}

			for _, platform := range targetPlatforms {
				arches, ok := bin.URLS[platform]
				if !ok {
					return fmt.Errorf(
						"binary '%s' missing platform '%s' in urls (required for architecture filtering)",
						bin.NAME, platform,
					)
				}

				for _, arch := range opts.SpecificArchBuilds {
					// URLs
					if _, ok := arches[arch]; !ok {
						return fmt.Errorf(
							"binary '%s' missing URL for platform '%s', architecture '%s'",
							bin.NAME, platform, arch,
						)
					}

					// SHA256
					if _, ok := bin.SHA256[platform][arch]; !ok {
						return fmt.Errorf(
							"binary '%s' missing SHA256 for platform '%s', architecture '%s'",
							bin.NAME, platform, arch,
						)
					}

					// PATTERNS are optional — skip
				}
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
