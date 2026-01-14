package arguments

import (
	"errors"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Validate the arguments with the config provided
func ValidateWithConfig(a *t.ArgOptions, c *t.Config) error {
	for _, p := range *c {
		if !slices.Contains(a.Packages, p.Name) {
			return errors.New("Packages passed is not found in the binman.yml file " + strings.Join(a.Packages, ", "))
		}
	}

	return nil
}
