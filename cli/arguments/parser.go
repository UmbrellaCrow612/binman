package arguments

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strings"

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

	platforms := flagSet.String("platforms", "", "Comma-separated list of platforms")
	architectures := flagSet.String("architectures", "", "Comma-separated list of architectures")

	err = flagSet.Parse(remainingArgs)
	if err != nil {
		return &options, err
	}

	options.Platforms = strings.Split(*architectures, ",")
	options.Platforms = strings.Split(*platforms, ",")

	return &options, nil
}
