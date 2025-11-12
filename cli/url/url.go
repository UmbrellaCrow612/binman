package url

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/yml"
)

var templateRegex = regexp.MustCompile(`\{\{(.+?)\}\}`)

func ResolveAllURLs(cfg *yml.Config) map[string]map[string]string {
	result := make(map[string]map[string]string)

	for _, bin := range cfg.Binaries {
		result[bin.Name] = make(map[string]string)

		for key, url := range bin.URL {
			// Find all {{key}} placeholders
			matches := templateRegex.FindAllStringSubmatch(url, -1)
			for _, match := range matches {
				if len(match) < 2 {
					continue
				}
				placeholder := match[0] // "{{version}}"
				fieldKey := match[1]    // "version"

				value, ok := bin.Extra[fieldKey]
				if !ok || value == "" {
					printer.ExitError(fmt.Sprintf("Binary '%s' URL '%s' requires '%s' but it is missing in extra fields", bin.Name, key, fieldKey))
				}

				url = strings.ReplaceAll(url, placeholder, value)
			}

			result[bin.Name][key] = url
		}
	}

	printer.PrintSuccess("All URLs resolved successfully")
	return result
}