package args

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/printer"
)

// List of options passed from the CLI mapped to fields
type Options struct {
	// The path to run the cli logic in
	Path string

	// Path to the binman.yml config file
	PathToFile string

	// Build only specific platforms keys like [linux, windows] - defaults to empty
	SpecificPlatformBuilds []string

	// If it should run the clean logic i.e via the pattern regex for the platform
	NoClean bool
}

// Parse args passed to the cli and get the options
func Parse() *Options {
	options := &Options{
		Path:                   "",
		PathToFile:             "",
		SpecificPlatformBuilds: []string{},
		NoClean:                false,
	}
	setOptions(options)

	return options
}

// Sets options values from args array
func setOptions(options *Options) {
	args := os.Args[1:]

	if len(args) == 0 {
		printer.ExitError("Missing path argument. Usage: binman <path> [..flags..]")
	}

	inputPath := args[0]
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		printer.ExitError("Failed to resolve path: " + err.Error())
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		printer.ExitError("Path does not exist: " + absPath)
	}
	if err == nil && !info.IsDir() {
		printer.ExitError("Provided path is not a directory: " + absPath)
	}

	options.Path = absPath
	printer.PrintSuccess("Resolved path: " + absPath)

	configPath := filepath.Join(absPath, "binman.yml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		printer.ExitError("Missing required config file: " + configPath)
	}

	options.PathToFile = configPath
	printer.PrintSuccess("Found config file: " + configPath)

	for _, arg := range args[1:] {
		switch {
		case strings.HasPrefix(arg, "--platforms="):
			value := strings.TrimPrefix(arg, "--platforms=")
			options.SpecificPlatformBuilds = strings.Split(value, ",")
			printer.PrintSuccess("Target platforms: " + strings.Join(options.SpecificPlatformBuilds, ", "))
		case arg == "--no-clean":
			options.NoClean = true
		default:
			printer.ExitError("Unknown flag: " + arg)
		}
	}
}
