package url

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
)

var templateRegex = regexp.MustCompile(`\{\{(.+?)\}\}`)

// ResolveAllURLs replaces {{placeholders}} in each Binary's URLs using Extra fields
func ResolveAllURLs(cfg *shared.Config) {
	for _, bin := range cfg.Binaries {
		for key, url := range bin.URL {
			// Find all {{key}} placeholders
			matches := templateRegex.FindAllStringSubmatch(url, -1)
			for _, match := range matches {
				if len(match) < 2 {
					continue
				}
				placeholder := match[0] // "{{version}}"
				fieldKey := match[1]    // "version"

				// Look up value in Extra
				value := ""
				if v, ok := bin.Extra[fieldKey]; ok {
					if s, ok := v.(string); ok {
						value = s
					} else {
						printer.ExitError(fmt.Sprintf(
							"Binary '%s' extra field '%s' is not a string",
							bin.Name, fieldKey,
						))
					}
				}

				if value == "" {
					printer.ExitError(fmt.Sprintf(
						"Binary '%s' URL '%s' requires '{{%s}}' but it is missing in Extra fields",
						bin.Name, key, fieldKey,
					))
				}

				url = strings.ReplaceAll(url, placeholder, value)
			}

			// Update the URL in-place
			bin.URL[key] = url
		}
	}

	printer.PrintSuccess("All URLs resolved successfully")
}
