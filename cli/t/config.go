package t

// Config is the root level of your YAML file.
// It is a slice of Packages.
//
// YAML Example:
// - name: ripgrep
//   platforms: { ... }
type Config []Package

// Package represents a software tool and its distribution data across OS/Arch.
type Package struct {
	// Name is the human-readable name of the package.
	Name string `yaml:"name"`

	// Platforms maps an Operating System (e.g., "linux", "windows", "darwin")
	// to a collection of available CPU architectures.
	//
	// YAML Example:
	// platforms:
	//   linux:
	//     x64: { ... }
	Platforms map[string]map[string]Asset `yaml:"platforms"`
}

// Asset contains the specific metadata required to download and verify
// a binary for a specific OS and Architecture combination.
//
// YAML Example:
// x64:
//   url: https://github.com/...
//   sha256: "234..."
//   pattern: "bin/rg"
type Asset struct {
	// URL is the direct download link for the package archive or binary.
	URL string `yaml:"url"`

	// SHA256 is the checksum used to verify that the file hasn't been tampered with.
	SHA256 string `yaml:"sha256"`

	// Pattern is a string or regex used to find the specific binary
	// within a downloaded .tar.gz or .zip file.
	Pattern string `yaml:"pattern"`
}
