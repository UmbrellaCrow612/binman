package arguments

import (
	"errors"
	"flag"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/global"
	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Parses the arguments / args passed to the CLI
func Parse() (*t.ArgOptions, error) {
	options := t.ArgOptions{}

	args := os.Args[1:]

	if len(args) < 1 {
		return &options, errors.New("Path not passed as first argument")
	}

	absBasePath, err := filepath.Abs(args[0])
	if err != nil {
		return &options, err
	}

	basePathInfo, err := os.Stat(absBasePath)
	if err != nil {
		return &options, err
	}
	if !basePathInfo.IsDir() {
		return &options, errors.New("Base path must point to a directory")
	}
	options.BasePath = absBasePath

	remainingArgs := args[1:] // these are all the flags

	flagSet := flag.NewFlagSet("binman", flag.ContinueOnError)

	platforms := flagSet.String("platforms", "", "Comma-separated list of platforms for example: linux,windows,darwin")
	architectures := flagSet.String("architectures", "", "Comma-separated list of architectures for example x64,arm64")
	verbose := flagSet.Bool("verbose", true, "Bool flag to produce logs or not during runtime")

	err = flagSet.Parse(remainingArgs)
	if err != nil {
		return &options, err
	}

	splitAndFilter := func(s string) []string {
		var result []string
		for item := range strings.SplitSeq(s, ",") {
			if trimmed := strings.TrimSpace(item); trimmed != "" {
				result = append(result, trimmed)
			}
		}
		return result
	}

	options.Platforms = splitAndFilter(*platforms)
	options.Architectures = splitAndFilter(*architectures)
	global.Verbose = *verbose

	for _, p := range options.Platforms {
		if !slices.Contains(global.ValidPlatforms, p) {
			return &options, errors.New("Platform passed is not valid: " + p + " - valid platforms are " + strings.Join(global.ValidPlatforms, ", "))
		}
	}

	for _, a := range options.Architectures {
		if !slices.Contains(global.ValidArchitectures, a) {
			return &options, errors.New("Architecture passed is not valid: " + a + " - valid architectures are " + strings.Join(global.ValidArchitectures, ", "))
		}
	}

	binmanYmlPath, err := filepath.Abs(path.Join(options.BasePath, "binman.yml"))
	if err != nil {
		return &options, err
	}

	binmanYmlFileInfo, err := os.Stat(binmanYmlPath)
	if err != nil {
		return &options, err
	}
	if binmanYmlFileInfo.IsDir() {
		return &options, errors.New("binman.yml cannot be a directory or a directory exists with the name binman.yml")
	}
	options.ConfigPath = binmanYmlPath

	return &options, nil
}
