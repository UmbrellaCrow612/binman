package t

// Config represents the top-level list of all software packages.
// YAML Example:
// - name: ripgrep
// - name: fd
type Config []Package

// Package defines a specific tool and the different ways it can be downloaded.
// YAML Example:
// name: ripgrep
// platforms: [...]
type Package struct {
	// Name is the identifier of the package (e.g., "ripgrep").
	Name string `yaml:"name"`

	// Platforms is a list because of the dash ('-') before the OS name in your YAML.
	// It contains maps where the key is the OS (e.g., "linux", "darwin").
	Platforms []Platform `yaml:"platforms"`
}

// Platform maps an Operating System name to a list of available CPU architectures.
// YAML Example:
// - linux:
//     - x64: [...]
type Platform map[string][]Architecture

// Architecture maps a specific CPU design to its download and verification metadata.
// YAML Example:
// - x64:
//     url: "..."
type Architecture map[string]Asset

// Asset contains the final technical details needed to download and verify a binary.
// YAML Example:
// url: https://github.com/...
// sha256: "sha:234..."
// pattern: "look for..."
type Asset struct {
	// URL is the direct link to the compressed archive or binary.
	URL string `yaml:"url"`

	// SHA256 is the checksum used to verify the integrity of the downloaded file.
	SHA256 string `yaml:"sha256"`

	// Pattern is a regex or string used to identify the correct file within an archive.
	Pattern string `yaml:"pattern"`
}
