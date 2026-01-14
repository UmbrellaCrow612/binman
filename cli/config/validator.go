package config

import (
	"errors"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/global"
	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Vaidate a parsed yml file config
func Validate(conf *t.Config) error {
	for _, pack := range *conf {
		err := validatePackage(pack)
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate a config package
func validatePackage(p t.Package) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("Package name cannot be a empty string")
	}

	for platform := range p.Platforms {
		if !slices.Contains(global.ValidPlatforms, platform) {
			return errors.New("Platform not valid: " + platform + ". Valid platforms are: " + strings.Join(global.ValidPlatforms, ", "))
		}
	}

	return nil
}
