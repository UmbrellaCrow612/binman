package t

// Config represents the entire binman.yml file.
//
// Example binman.yml:
//
// - name: ripgrep
//   linux:
//     x64:
//       url: "https://example.com/rg-linux-x64.tar.gz"
//       sha256: "abc123"
//       pattern: "rg-*"
//   windows:
//     x64:
//       url: "https://example.com/rg-win-x64.zip"
//       sha256: "def456"
//       pattern: "rg.exe"
//
type Config []Entry

// Entry represents a single binary definition.
//
// Example:
//
// - name: ripgrep
//   linux:
//     x64:
//       url: "..."
//       sha256: "..."
//       pattern: "..."
//
type Entry struct {
	// Name of the binary/tool
	//
	// name: ripgrep
	Name string `yaml:"name"`

	// Platforms defines per-platform configurations.
	// The key is the platform name (linux, windows, darwin, etc).
	//
	// linux:
	//   x64:
	//     url: "..."
	//
	// windows:
	//   x64:
	//     url: "..."
	//
	// NOTE:
	// This is inlined so platforms appear at the same level as `name`
	Platforms map[string]PlatformDefinition `yaml:",inline"`
}

// PlatformDefinition represents all architectures for a platform.
//
// Example:
//
// linux:
//   x64:
//     url: "..."
//     sha256: "..."
//     pattern: "rg-*"
//   arm64:
//     url: "..."
//     sha256: "..."
//     pattern: "rg-*"
//
type PlatformDefinition map[string]ArchitectureDefinition

// ArchitectureDefinition represents a single architecture download.
//
// Example:
//
// x64:
//   url: "https://example.com/rg-linux-x64.tar.gz"
//   sha256: "abc123"
//   pattern: "rg-*"
//
type ArchitectureDefinition struct {
	// Download URL for this platform + architecture
	//
	// url: "https://example.com/rg-linux-x64.tar.gz"
	URL string `yaml:"url"`

	// SHA256 checksum for verifying the downloaded binary
	//
	// sha256: "abc123"
	SHA256 string `yaml:"sha256"`

	// Filename pattern to match inside extracted archives
	//
	// pattern: "rg-*"
	Pattern string `yaml:"pattern"`
}
