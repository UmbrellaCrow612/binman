package yml

import (
	"os"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"github.com/UmbrellaCrow612/binman/cli/shared"
	"gopkg.in/yaml.v2"
)

// parse reads the YAML file from opts.PathToFile and returns the parsed Config
func Parse(opts *args.Options) *shared.Config {
	if opts.PathToFile == "" {
		printer.ExitError("PathToFile is empty")
	}

	data, err := os.ReadFile(opts.PathToFile)
	if err != nil {
		printer.ExitError("Failed to read file: " + err.Error())
	}

	var cfg shared.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		printer.ExitError("Failed to parse YAML: " + err.Error())
	}

	err = cfg.Validate()
	if err != nil {
		printer.ExitError("Failed to parse YAML: " + err.Error())
	}

	printer.PrintSuccess("YAML file parsed successfully")

	return &cfg
}
