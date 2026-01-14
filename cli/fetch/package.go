package fetch

import (
	"path/filepath"

	"github.com/UmbrellaCrow612/binman/cli/t"
)

// Fetch a package into ./downloads/packname/platform/arch/...
// Using the options defined
func Get(p *t.Package, c *t.ArgOptions) error {
	_, err := filepath.Abs(filepath.Join(c.BasePath, "downloads"))
	if err != nil {
		return err
	}

	return nil
}
