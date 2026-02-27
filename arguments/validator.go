package arguments

import (
	"errors"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/t"
)

// Validate the arguments with the config provided
func ValidateWithConfig(a *t.ArgOptions, c *t.Config) error {
	binmanPackages := []string{}
	for _, p := range *c {
		binmanPackages = append(binmanPackages, p.Name)
	}

	for _, p := range a.Packages {
		if len(a.Packages) > 0 && !slices.Contains(binmanPackages, p) {
			return errors.New("Package passed is not found in the binman.yml file: " + p + " - binman packages " + strings.Join(binmanPackages, ", "))
		}
	}

	return nil
}
