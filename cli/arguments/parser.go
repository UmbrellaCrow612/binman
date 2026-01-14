package arguments

import (
	"errors"
	"os"

	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Parses the arguments / args passed to the CLI
func Parse() (*t.ArgOptions, error) {
	options := t.ArgOptions{}

	args := os.Args[1:]

	if len(args) < 1 {
		return &options, errors.New("Path not passed as first argument")
	}

	return &options, nil
}
