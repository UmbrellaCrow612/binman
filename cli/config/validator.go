package config

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/global"
	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Vaidate a parsed yml file config
func Validate(conf *t.Config) error {
	for _, pack := range *conf {
		err := validatePackage(&pack)
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate a config package
func validatePackage(p *t.Package) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("Package name cannot be a empty string")
	}

	for platform, archMap := range p.Platforms {
		if !slices.Contains(global.ValidPlatforms, platform) {
			return errors.New("Platform not valid: " + platform + ". Valid platforms are: " + strings.Join(global.ValidPlatforms, ", "))
		}

		for arch, asset := range archMap {
			if !slices.Contains(global.ValidArchitectures, arch) {
				return errors.New("Architecture not valid: " + arch + ". Valid architectures are: " + strings.Join(global.ValidArchitectures, ", "))
			}

			if strings.TrimSpace(asset.Pattern) == "" {
				return errors.New("Pattern must be defined for " + platform + " " + arch)
			}

			_, err := regexp.Compile(asset.Pattern)
			if err != nil {
				return err
			}

			if strings.TrimSpace(asset.SHA256) == "" {
				return errors.New("SHA256 must be defined for " + platform + " " + arch)
			}

			const sha256Prefix = "sha256:"
			asset.SHA256 = strings.TrimPrefix(asset.SHA256, sha256Prefix)
		}
	}

	return nil
}
